package processes

import (
	"errors"
	"os"

	"go.mondoo.io/mondoo/motor"
)

type OSProcess struct {
	Pid        int64
	Command    string
	Executable string
	State      string
	Uid        int64
}

type OSProcessManager interface {
	Name() string
	Exists(pid int64) (bool, error)
	Process(pid int64) (*OSProcess, error)
	List() ([]*OSProcess, error)
}

func ResolveManager(motor *motor.Motor) (OSProcessManager, error) {
	var pm OSProcessManager

	platform, err := motor.Platform()
	if err != nil {
		return nil, err
	}

	switch {
	case platform.IsFamily("linux") && os.Getenv("MONDOO_PROCFS") == "on":
		pm = &LinuxProcManager{motor: motor}
	case platform.IsFamily("unix"):
		pm = &UnixProcessManager{motor: motor, platform: platform}
	case platform.IsFamily("windows"):
		pm = &WindowsProcessManager{motor: motor}
	default:
		return nil, errors.New("could not detect suitable process manager for platform: " + platform.Name)
	}

	return pm, nil
}
