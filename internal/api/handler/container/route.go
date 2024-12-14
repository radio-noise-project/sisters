package container

import (
	"io"
	"log/slog"
	"os"
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

	filePath := "archive.tar.gz"
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	slog.Info("Upload completed")
	return stream.SendAndClose(&StatusResponse{Status: "Success"})
}
