package handler

import (
	"bufio"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
	"github.com/radio-noise-project/sisters/pkg/runtime"
)

func OutputSistersVersion(c echo.Context) error {
	info := VersionInformationSisters.getVersionInformation(VersionInformationSisters{})
	return c.JSON(http.StatusOK, info)
}

func (VersionInformationSisters) getVersionInformation() *VersionInformationSisters {
	codeName, version, builtDate := getSistersVersion()
	golangVerson := getGolangVersion()
	dockerEngineVersion := runtime.DockerEngineVersion()
	revisionHash := getGitCommitHash()
	os, arch := getOsArchVersion()
	info := &VersionInformationSisters{
		CodeName:            codeName,
		Version:             version,
		GolangVersion:       golangVerson,
		DockerEngineVersion: dockerEngineVersion,
		BuiltGitCommitHash:  revisionHash,
		BuiltDate:           strToTime(builtDate),
		Os:                  os,
		Arch:                arch,
	}
	return info
}

func getSistersVersion() (string, string, string) {
	fp, err := os.Open("../../VERSION")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	var str [3]string
	var i = 0
	for scanner.Scan() {
		str[i] = scanner.Text()
		i++
	}

	codeName := str[0]
	version := str[1]
	builtDate := str[2]
	return codeName, version, builtDate
}

func strToTime(t string) time.Time {
	prasedTime, err := time.Parse("2006-01-02T15:04:05Z07:00", t)
	if err != nil {
		panic(err)
	}

	return prasedTime
}
