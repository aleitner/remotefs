package server

import (
	"net"

	rfspb "remotefs/pkg/protobuf"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type GrpcRemoteFileServer struct {
	directory string
	logger *log.Logger
}

func NewGrpcRFS(logger *log.Logger, dir string) (*GrpcRemoteFileServer, error) {
	return &GrpcRemoteFileServer{
		directory: dir,
		logger:    logger,
	}, nil
}

func(rfs *GrpcRemoteFileServer) Serve(l net.Listener) {
	grpcS := grpc.NewServer()
	rfspb.RegisterFTransferServer(grpcS, rfs)

	grpcS.Serve(l)
}

func(rfs *GrpcRemoteFileServer) Upload(server rfspb.FTransfer_UploadServer) error {
	return nil
}