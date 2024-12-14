package container

import (
	"io"
	"log/slog"

	"github.com/radio-noise-project/sisters/pkg/api/container"
)

type Server struct {
	UnimplementedContainerServiceServer
}

func (s *Server) Upload(stream ContainerService_UploadServer) error {
	slog.Info("Upload called")
	var data []byte

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		data = append(data, req.GetArchive()...)
		slog.Info("Received chunk of data")
	}

	container.ContainerCreateAndStart(data)
	return stream.SendAndClose(&StatusResponse{Status: "Success"})
}
