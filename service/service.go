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
	uploadservice := &uploadService{}
	upload.RegisterUploadHandlerServer(gserver, uploadservice)
	reflection.Register(gserver)
}

type uploadService struct{}

func (s *uploadService) Upload(stream upload.UploadHandler_UploadServer) error {
	err := os.MkdirAll("Sample", 0777)
	if err != nil {
		return err
	}
	file, err := os.Create(filepath.Join("resource", "tmp.mp4"))
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
		_, err = file.Write(resp.VideoData)
		if err != nil {
			return err
		}

	}
	err = stream.SendAndClose(&upload.UploadReply{UploadStatus: "OK"})
	if err != nil {

		return err
	}
	return nil

}
