package runtime

import (
	"context"
	"log/slog"

	"github.com/radio-noise-project/sisters/pkg/api/runtime"
)

type Server struct {
}

func (s *Server) Version(ctx context.Context) (*VersionResponse, error) {
	slog.Info("GetVersion called")
	codeName, version, golanVersion, dockerEngineVersion, builtGitCommitHash, builtDate, os, arch := runtime.GetVersion()

	return &VersionResponse{
		CodeName:            codeName,
		Version:             version,
		GolangVersion:       golanVersion,
		DockerEngineVersion: dockerEngineVersion,
		BuiltGitcommitHash:  builtGitCommitHash,
		BuiltDate:           builtDate,
		Os:                  os,
		Arch:                arch,
	}, nil
}
