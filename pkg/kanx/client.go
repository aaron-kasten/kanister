package kanx

import (
	"context"
	"io"
	"net"
	"net/url"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func unixDialer(ctx context.Context, addr string) (net.Conn, error) {
	return net.Dial("unix", addr)
}

func newGRPCConnection(addr string) (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithContextDialer(unixDialer))
	// Add passthrough scheme if there is no scheme defined in the address
	u, err := url.Parse(addr)
	if err == nil && u.Scheme == "" {
		addr = "passthrough:///" + addr
	}
	return grpc.NewClient(addr, opts...)
}

func CreateProcess(ctx context.Context, addr string, name string, args []string) (*Process, error) {
	conn, err := newGRPCConnection(addr)
	if err != nil {
		return nil, err
	}
	defer conn.Close() //nolint:errcheck
	c := NewProcessServiceClient(conn)
	in := &CreateProcessRequest{
		Name: name,
		Args: args,
	}
	return c.CreateProcess(ctx, in)
}

func ListProcesses(ctx context.Context, addr string) ([]*Process, error) {
	conn, err := newGRPCConnection(addr)
	if err != nil {
		return nil, err
	}
	defer conn.Close() //nolint:errcheck
	c := NewProcessServiceClient(conn)
	lpc, err := c.ListProcesses(ctx, &ListProcessesRequest{})
	if err != nil {
		return nil, err
	}
	ps := []*Process{}
	for {
		p, err := lpc.Recv()
		switch {
		case err == io.EOF:
			return ps, nil
		case err != nil:
			return nil, err
		}
		ps = append(ps, p)
	}
}

func WaitProcess(ctx context.Context, addr string, pid int64) (*Process, error) {
	conn, err := newGRPCConnection(addr)
	if err != nil {
		return nil, err
	}
	defer conn.Close() //nolint:errcheck
	c := NewProcessServiceClient(conn)
	wpc, err := c.WaitProcess(ctx, &ProcessPidRequest{
		Pid: pid,
	})
	if err != nil {
		return nil, err
	}
	return wpc, nil
}

func SignalProcess(ctx context.Context, addr string, pid int64, sig int32) (*Process, error) {
	conn, err := newGRPCConnection(addr)
	if err != nil {
		return nil, err
	}
	defer conn.Close() //nolint:errcheck
	c := NewProcessServiceClient(conn)
	wpc, err := c.SignalProcess(ctx, &SignalProcessRequest{
		Pid:    pid,
		Signal: sig,
	})
	if err != nil {
		return nil, err
	}
	return wpc, nil
}

func GetProcess(ctx context.Context, addr string, pid int64) (*Process, error) {
	conn, err := newGRPCConnection(addr)
	if err != nil {
		return nil, err
	}
	defer conn.Close() //nolint:errcheck
	c := NewProcessServiceClient(conn)
	wpc, err := c.GetProcess(ctx, &ProcessPidRequest{
		Pid: pid,
	})
	if err != nil {
		return nil, err
	}
	return wpc, nil
}

func RemoveProcess(ctx context.Context, addr string, pid int64) (*Process, error) {
	conn, err := newGRPCConnection(addr)
	if err != nil {
		return nil, err
	}
	defer conn.Close() //nolint:errcheck
	c := NewProcessServiceClient(conn)
	wpc, err := c.RemoveProcess(ctx, &ProcessPidRequest{
		Pid: pid,
	})
	if err != nil {
		return nil, err
	}
	return wpc, nil
}

func Stdout(ctx context.Context, addr string, pid int64, out io.Writer) error {
	conn, err := newGRPCConnection(addr)
	if err != nil {
		return err
	}
	defer conn.Close() //nolint:errcheck
	c := NewProcessServiceClient(conn)
	in := &ProcessPidRequest{
		Pid: pid,
	}
	poc, err := c.Stdout(ctx, in)
	if err != nil {
		return err
	}
	return stdoutOutput(ctx, out, poc)
}

func Stderr(ctx context.Context, addr string, pid int64, out io.Writer) error {
	conn, err := newGRPCConnection(addr)
	if err != nil {
		return err
	}
	defer conn.Close() //nolint:errcheck
	c := NewProcessServiceClient(conn)
	in := &ProcessPidRequest{
		Pid: pid,
	}
	poc, err := c.Stderr(ctx, in)
	if err != nil {
		return err
	}
	return stderrOutput(ctx, out, poc)
}

type stdoutRecver interface {
	Recv() (*StdoutOutput, error)
}

type stderrRecver interface {
	Recv() (*StderrOutput, error)
}

func stderrOutput(ctx context.Context, out io.Writer, sc stderrRecver) error {
	for {
		p, err := sc.Recv()
		switch {
		case err == io.EOF:
			return nil
		case err != nil:
			return err
		}
		_, err = out.Write([]byte(p.Stream))
		if err != nil {
			return err
		}
	}
}

func stdoutOutput(ctx context.Context, out io.Writer, sc stdoutRecver) error {
	for {
		p, err := sc.Recv()
		switch {
		case err == io.EOF:
			return nil
		case err != nil:
			return err
		}
		_, err = out.Write([]byte(p.Stream))
		if err != nil {
			return err
		}
	}
}
