package metrics

import "os/exec"

// GetenvFn is used to mock os.Getenv() for testing only
type GetenvFn func(key string) string

// NewTestMetrics create a full mocking testing metrics element.
// This is only a mock for testing, not for real use
func NewTestMetrics(root string,
	cmdGPU *exec.Cmd, cmdCPU *exec.Cmd, cmdScreen *exec.Cmd, cmdPartition *exec.Cmd,
	cmdArch *exec.Cmd, cmdLibc6 *exec.Cmd, cmdHwCap *exec.Cmd, getenv GetenvFn) Metrics {
	// do not use helper as in _test.go package
	return Metrics{
		root:          root,
		gpuInfoCmd:    cmdGPU,
		cpuInfoCmd:    cmdCPU,
		screenInfoCmd: cmdScreen,
		spaceInfoCmd:  cmdPartition,
		archCmd:       cmdArch,
		libc6Cmd:      cmdLibc6,
		hwCapCmd:      cmdHwCap,
		getenv:        getenv,
	}
}
