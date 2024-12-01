package runtime

import (
	"github.com/radio-noise-project/sisters/internal/api/pb"
	"google.golang.org/grpc"
)

func Handler(s *grpc.Server) {
	pb.RegisterRuntimeServiceServer(s, server())
}
