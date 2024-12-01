package runtime

import (
	"context"

	"github.com/radio-noise-project/sisters/internal/api/pb"
	"github.com/radio-noise-project/sisters/pkg/api/runtime"
)

type version struct {
	pb.UnimplementedRuntimeServiceServer
}

func (s *version) GetVersion(ctx context.Context, req *pb.VersionRequest) (*pb.VersionResponse, error) {
	codeName, version, golanVersion, dockerEngineVersion, builtGitCommitHash, builtDate, os, arch := runtime.GetVersion()

	return &pb.VersionResponse{
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

func server() *version {
	return &version{}
}
