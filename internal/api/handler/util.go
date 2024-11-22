package handler

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

func getGolangVersion() string {
	version := runtime.Version()
	return version
}

func getOsArchVersion() (string, string) {
	os := runtime.GOOS
	arch := runtime.GOARCH
	return os, arch
}

func getGitCommitHash() string {
	var hash string
	info, err := debug.ReadBuildInfo()
	if !err {
		fmt.Print("Nothing build information")
	}
	for _, s := range info.Settings {
		if s.Key == "vcs.revision" {
			hash = s.Value
		}
	}
	return hash
}
