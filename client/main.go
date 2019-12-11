package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"google.golang.org/grpc"

	upload "github.com/tetsuzawa/SOFA-Uplorder/proto/sofa"
)

func main() {
	connect, _ := grpc.Dial("localhost:8083", grpc.WithInsecure())

	defer connect.Close()
	uploadhalder := upload.NewUploadHandlerClient(connect)
	stream, err := uploadhalder.Upload(context.Background())
	if err != nil {
		fmt.Printf("failed to make stream: %v", err)
	}
	err = Upload(stream)
	if err != nil {
		fmt.Printf("failed to upload: %v", err)
	}
}

func Upload(stream upload.UploadHandler_UploadClient) error {
	file, err := os.Open("resource/wave.mp4")
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()
	buf := make([]byte, 1024)

	for {
		_, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("io error: %w", err)
		}
		err = stream.Send(&upload.UploadRequest{VideoData: buf})
		if err != nil {
			return fmt.Errorf("stream send error: %w", err)
		}
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		return fmt.Errorf("failed to CloseAndRecv: %w",err)
	}
	fmt.Println(resp.UploadStatus)
	return nil
}
