package kanx

import (
	"context"
	"io"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"github.com/kanisterio/errkit"
	"google.golang.org/grpc"

	"github.com/kanisterio/kanister/pkg/field"
	"github.com/kanisterio/kanister/pkg/log"
	"sync"
	"maps"
)

const (
	tailTickDuration  = 100 * time.Millisecond
	tempStdoutPattern = "kando.*.stdout"
	tempStderrPattern = "kando.*.stderr"
	streamBufferBytes = (4 * 1024 * 1024) / 2 // 4MiB max buffer size, remove to allow for JSON overhead
)

type processServiceServer struct {
	UnimplementedProcessServiceServer
	mu               sync.Mutex
	processByPid     map[int64]*process
	processes        []*process
	outputDir        string
	tailTickDuration time.Duration
}

type process struct {
	mu       sync.Mutex
	cmd      *exec.Cmd
	doneCh   chan struct{}
	stdout   *os.File
	stderr   *os.File
	exitCode int
	err      error
	fault    error
	cancel   context.CancelFunc
}

func newProcessServiceServer() *processServiceServer {
	return &processServiceServer{
		processByPid:     map[int64]*process{},
		tailTickDuration: tailTickDuration,
	}
}

func (s *processServiceServer) CreateProcess(_ context.Context, cpr *CreateProcessRequest) (*Process, error) {
	stdout, err := os.CreateTemp(s.outputDir, tempStdoutPattern)
	if err != nil {
		return nil, err
	}
	stderr, err := os.CreateTemp(s.outputDir, tempStderrPattern)
	if err != nil {
		return nil, err
	}
	// We use context.Background() here because the parameter ctx seems to get canceled when this returns.
	ctx, can := context.WithCancel(context.Background())
	cmd := exec.CommandContext(ctx, cpr.GetName(), cpr.GetArgs()...)
	p := &process{
		cmd:    cmd,
		doneCh: make(chan struct{}),
		stdout: stdout,
		stderr: stderr,
		cancel: can,
	}
	stdoutLogWriter := newLogWriter(log.Info(), os.Stdout)
	stderrLogWriter := newLogWriter(log.Info(), os.Stderr)
	cmd.Stdout = io.MultiWriter(p.stdout, stdoutLogWriter)
	cmd.Stderr = io.MultiWriter(p.stderr, stderrLogWriter)

	err = cmd.Start()
	if err != nil {
		return nil, err
	}
	s.mu.Lock()
	s.processByPid[int64(cmd.Process.Pid)] = p
	s.processes = append(s.processes, p)
	s.mu.Unlock()
	fields := field.M{"pid": cmd.Process.Pid, "stdout": stdout.Name(), "stderr": stderr.Name()}
	stdoutLogWriter.SetFields(fields)
	stderrLogWriter.SetFields(fields)
	log.Info().Print(processToProtoWithLock(p).String(), fields)
	go func() {
		err := p.cmd.Wait()
		p.err = err
		if exiterr, ok := err.(*exec.ExitError); ok {
			p.exitCode = exiterr.ExitCode()
		}
		err = stdout.Close()
		if err != nil {
			log.Error().WithError(err).Print("Failed to close stdout", fields)
		}
		err = stderr.Close()
		if err != nil {
			log.Error().WithError(err).Print("Failed to close stderr", fields)
		}
		can()
		close(p.doneCh)
		log.Info().Print(processToProtoWithLock(p).String())
	}()
	return &Process{
		Pid:   int64(cmd.Process.Pid),
		State: ProcessState_PROCESS_STATE_RUNNING,
	}, nil
}

func (s *processServiceServer) GetProcess(ctx context.Context, grp *ProcessPidRequest) (*Process, error) {
	s.mu.Lock()
	q, ok := s.processByPid[grp.GetPid()]
	s.mu.Unlock()
	if !ok {
		return nil, errkit.WithStack(errProcessNotFound)
	}
	ps := processToProtoWithLock(q)
	return ps, nil
}

func (s *processServiceServer) RemoveProcess(ctx context.Context, grp *ProcessPidRequest) (*Process, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	q, ok := s.processByPid[grp.GetPid()]
	if !ok {
		return nil, errkit.WithStack(errProcessNotFound)
	}
	q.mu.Lock()
	defer q.mu.Unlock()
	terminated := false
	select {
	case <-q.doneCh:
		terminated = true
	default:
	}
	if !terminated {
		return processToProto(q), errkit.WithStack(errProcessRunning)
	}
	var err error
	if q.stderr != nil {
		err = os.Remove(q.stderr.Name())
	}
	if err == nil {
		q.stderr = nil
	} else {
		q.fault = err
		return processToProto(q), err
	}
	if q.stdout != nil {
		err = os.Remove(q.stdout.Name())
	}
	if err == nil {
		q.stdout = nil
	} else {
		q.fault = err
		return processToProto(q), err
	}
	delete(s.processByPid, grp.GetPid())
	return processToProto(q), nil
}

func (s *processServiceServer) SignalProcess(ctx context.Context, grp *SignalProcessRequest) (*Process, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	q, ok := s.processByPid[grp.GetPid()]
	if !ok {
		return nil, errkit.WithStack(errProcessNotFound)
	}
	q.mu.Lock()
	defer q.mu.Unlock()
	// low level signal call
	syssig := syscall.Signal(grp.Signal)
	err := q.cmd.Process.Signal(syssig)
	if err != nil {
		q.fault = err
		return processToProto(q), err
	}
	return processToProto(q), nil
}

func (s *processServiceServer) WaitProcess(ctx context.Context, wpr *ProcessPidRequest) (*Process, error) {
	s.mu.Lock()
	q, ok := s.processByPid[wpr.GetPid()]
	s.mu.Unlock()
	if !ok {
		return nil, errkit.WithStack(errProcessNotFound)
	}
	select {
	case <-ctx.Done():
		return processToProtoWithLock(q), ctx.Err()
	case <-q.doneCh:
		return processToProtoWithLock(q), nil
	}
}

func (s *processServiceServer) ListProcesses(lpr *ListProcessesRequest, lps ProcessService_ListProcessesServer) error {
	s.mu.Lock()
	processes := maps.Clone(s.processByPid)
	s.mu.Unlock()
	for _, p := range processes {
		ps := processToProtoWithLock(p)
		err := lps.Send(ps)
		if err != nil {
			return err
		}
	}
	return nil
}

var (
	errProcessNotFound = errkit.NewSentinelErr("Process not found")
	errProcessRunning  = errkit.NewSentinelErr("Process running")
)

func (s *processServiceServer) Stdout(por *ProcessPidRequest, ss ProcessService_StdoutServer) error {
	s.mu.Lock()
	p, ok := s.processByPid[por.Pid]
	s.mu.Unlock()
	if !ok {
		return errkit.WithStack(errProcessNotFound)
	}
	fh, err := os.Open(p.stdout.Name())
	if err != nil {
		return err
	}
	err = s.streamStdoutOutput(ss, p, fh)
	if err == io.EOF {
		return nil
	}
	return err
}

func (s *processServiceServer) Stderr(por *ProcessPidRequest, ss ProcessService_StderrServer) error {
	s.mu.Lock()
	p, ok := s.processByPid[por.Pid]
	s.mu.Unlock()
	if !ok {
		return errkit.WithStack(errProcessNotFound)
	}
	fh, err := os.Open(p.stderr.Name())
	if err != nil {
		return err
	}
	err = s.streamStderrOutput(ss, p, fh)
	if err == io.EOF {
		return nil
	}
	return err
}

type stdoutSender interface {
	Send(*StdoutOutput) error
}

type stderrSender interface {
	Send(*StderrOutput) error
}

func (s *processServiceServer) streamStdoutOutput(ss stdoutSender, p *process, fh *os.File) error {
	buf := [streamBufferBytes]byte{} // 4MiB is the max size of a GRPC request
	t := time.NewTicker(s.tailTickDuration)
	for {
		n, err := fh.Read(buf[:])
		switch {
		case err == io.EOF && n == 0:
			select {
			case <-p.doneCh:
				return nil
			default:
			}
			<-t.C
			continue
		case err != nil:
			return err
		}
		o := &StdoutOutput{Stream: string(buf[:n])}
		err = ss.Send(o)
		if err != nil {
			return err
		}
	}
}

func (s *processServiceServer) streamStderrOutput(ss stderrSender, p *process, fh *os.File) error {
	buf := [streamBufferBytes]byte{} // 4MiB is the max size of a GRPC request
	t := time.NewTicker(s.tailTickDuration)
	for {
		n, err := fh.Read(buf[:])
		switch {
		case err == io.EOF && n == 0:
			select {
			case <-p.doneCh:
				return nil
			default:
			}
			<-t.C
			continue
		case err != nil:
			return err
		}
		o := &StderrOutput{Stream: string(buf[:n])}
		err = ss.Send(o)
		if err != nil {
			return err
		}
	}
}

func processToProto(p *process) *Process {
	ps := &Process{
		Pid: int64(p.cmd.Process.Pid),
	}
	select {
	case <-p.doneCh:
		ps.State = ProcessState_PROCESS_STATE_SUCCEEDED
		if p.err != nil {
			ps.State = ProcessState_PROCESS_STATE_FAILED
			ps.ExitErr = p.err.Error()
			ps.ExitCode = int64(p.exitCode)
		}
	default:
		ps.State = ProcessState_PROCESS_STATE_RUNNING
	}
	return ps
}

func processToProtoWithLock(p *process) *Process {
	p.mu.Lock()
	defer p.mu.Unlock()
	return processToProto(p)
}

type Server struct {
	grpcs *grpc.Server
	pss   *processServiceServer
}

func NewServer() *Server {
	var opts []grpc.ServerOption
	return &Server{
		grpcs: grpc.NewServer(opts...),
		pss:   newProcessServiceServer(),
	}
}

func (s *Server) Serve(ctx context.Context, addr string) error {
	ctx, can := signal.NotifyContext(ctx, syscall.SIGTERM, syscall.SIGINT)
	defer can()
	go func() {
		<-ctx.Done()
		if err := ctx.Err(); err == context.Canceled {
			log.Info().Print("Gracefully stopping. Parent context canceled")
		} else {
			log.Info().WithError(err).Print("Gracefully stopping.")
		}
		s.grpcs.GracefulStop()
	}()
	RegisterProcessServiceServer(s.grpcs, s.pss)
	lis, err := net.Listen("unix", addr)
	if err != nil {
		return err
	}
	log.Info().Print("Listening on socket", field.M{"address": lis.Addr()})
	defer os.Remove(addr) //nolint:errcheck
	return s.grpcs.Serve(lis)
}
