# Release Notes

## 0.113.0

## New Features

<!-- releasenotes/notes/pre-release-0.113.0-6e4df1d2b04b3ca7.yaml @ None -->
* Allow signals to be sent to `kando` run processes

<!-- releasenotes/notes/pre-release-0.113.0-6e4df1d2b04b3ca7.yaml @ None -->
* Execute and follow output of `kando` run processes

<!-- releasenotes/notes/rds-credentials-1fa9817a21a2d80a.yaml @ b'c4534cdbb7167c6f854c4d7915dd22483f9486f9' -->
* Enable RDS functions to accept AWS credentials using a Secret or ServiceAccount.

## Bug Fixes

<!-- releasenotes/notes/pre-release-0.113.0-6e4df1d2b04b3ca7.yaml @ None -->
* The Kopia snapshot command output parser now skips the ignored and fatal error counts

<!-- releasenotes/notes/pre-release-0.113.0-6e4df1d2b04b3ca7.yaml @ None -->
* Set default namespace and serviceaccount for MultiContainerRun pods

## Upgrade Notes

<!-- releasenotes/notes/pre-release-0.113.0-6e4df1d2b04b3ca7.yaml @ None -->
* Upgrade to K8s 1.31 API

<!-- releasenotes/notes/pre-release-0.113.0-6e4df1d2b04b3ca7.yaml @ None -->
* Bump Kopia dependecy to git SHA 1bceb71

## Deprecations

<!-- releasenotes/notes/pre-release-0.113.0-6e4df1d2b04b3ca7.yaml @ None -->
* K8s VolumeSnapshot is now GA, remove support for beta and alpha APIs

## Other Notes

<!-- releasenotes/notes/pre-release-0.113.0-6e4df1d2b04b3ca7.yaml @ None -->
* Change `TIMEOUT_WORKER_POD_READY` environment variable to `KANISTER_POD_READY_WAIT_TIMEOUT`

<!-- releasenotes/notes/pre-release-0.113.0-6e4df1d2b04b3ca7.yaml @ None -->
* Errors are now handled with [https://github.com/kanisterio/errkit](https://github.com/kanisterio/errkit) across the board
