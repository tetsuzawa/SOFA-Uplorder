package service

import (
	"io"
	"os"
	"path/filepath"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	upload "github.com/tetsuzawa/SOFA-Uplorder/proto/sofa"
)

func NewUploadServer(gserver *grpc.Server) {
	uploadserver := &server{}
	upload.RegisterUploadHandlerServer(gserver, uploadserver)
	reflection.Register(gserver)
}

type server struct{}

func (s *server) Upload(stream upload.UploadHandler_UploadServer) error {
	err := os.MkdirAll("Sample", 0777)
	if err != nil {
		return err
	}
	file, err := os.Create(filepath.Join("Sample", "tmp.mp4"))
	defer file.Close()
	if err != nil {
		return err
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		file.Write(resp.VideoData)
	}
	err = stream.SendAndClose(&upload.UploadReply{UploadStatus: "OK"})
	if err != nil {

		return err
	}
	return nil

}
