package server

import (
	"github.com/radio-noise-project/sisters/internal/api/handler/runtime"
	"google.golang.org/grpc"
)

func router(s *grpc.Server) {
	runtime.Handler(s)
}
