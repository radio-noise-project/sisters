package server

import (
	"context"
	"log"
	"log/slog"
	"net"
	"os"
	"os/signal"

	"github.com/radio-noise-project/sisters/internal/api/handler/container"
	"github.com/radio-noise-project/sisters/internal/api/handler/runtime"
	runtimePkg "github.com/radio-noise-project/sisters/pkg/api/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

func Start() {
	slog.Info("Server Started")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	runtime.RegisterRuntimeServiceServer(grpcServer, &Server{})
	container.RegisterContainerServiceServer(grpcServer, &container.Server{})
	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	grpcServer.GracefulStop()
}

type Server struct {
	runtime.UnimplementedRuntimeServiceServer
}

func (s *Server) Version(ctx context.Context, empty *emptypb.Empty) (*runtime.VersionResponse, error) {
	codeName, version, golanVersion, dockerEngineVersion, builtGitCommitHash, builtDate, os, arch := runtimePkg.GetVersion()

	return &runtime.VersionResponse{
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
