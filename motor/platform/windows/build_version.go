package windows

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"strconv"

	"go.mondoo.com/cnquery/motor/providers/os"
	"go.mondoo.com/cnquery/motor/providers/os/powershell"
)

var WinBuildVersionRegex = regexp.MustCompile(`^(\d+)(?:\.(\d+)){0,1}`)

type BuildVersion struct {
	Build int
	UBR   int
}

// WindowsVersion parses the version generated by motor
// It uses the format: major.minor.build.ubr
func Version(ver string) (*BuildVersion, error) {
	// check that we go a valid windows version
	m := WinBuildVersionRegex.FindStringSubmatch(ver)
	if len(m) < 2 {
		return nil, fmt.Errorf("the requested windows version is not supported: %s", ver)
	}

	build, err := strconv.Atoi(m[1])
	if err != nil {
		return nil, err
	}

	var ubr int
	if len(m) == 3 && len(m[2]) > 0 {
		ubr, err = strconv.Atoi(m[2])
		if err != nil {
			return nil, err
		}
	}

	return &BuildVersion{
		Build: build,
		UBR:   ubr,
	}, nil
}

func (b BuildVersion) OSBuild() string {
	if b.UBR > 0 {
		return fmt.Sprintf("%d.%d", b.Build, b.UBR)
	} else {
		return fmt.Sprintf("%d", b.Build)
	}
}

type WindowsCurrentVersion struct {
	CurrentBuild string `json:"CurrentBuild"`
	EditionID    string `json:"EditionID"`
	ReleaseId    string `json:"ReleaseId"`
	// Update Build Revision
	UBR int `json:"UBR"`
}

func ParseWinRegistryCurrentVersion(r io.Reader) (*WindowsCurrentVersion, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var winCurrentVersion WindowsCurrentVersion
	err = json.Unmarshal(data, &winCurrentVersion)
	if err != nil {
		return nil, err
	}

	return &winCurrentVersion, nil
}

// powershellGetWindowsOSBuild runs a powershell script to retrieve the current version from windows
func powershellGetWindowsOSBuild(t os.OperatingSystemProvider) (*WindowsCurrentVersion, error) {
	pscommand := "Get-ItemProperty -Path 'HKLM:\\SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion' -Name CurrentBuild, UBR, EditionID | ConvertTo-Json"
	cmd, err := t.RunCommand(powershell.Wrap(pscommand))
	if err != nil {
		return nil, err
	}
	return ParseWinRegistryCurrentVersion(cmd.Stdout)
}
