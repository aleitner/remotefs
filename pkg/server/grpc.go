package server

import (
	"net"

	log "github.com/sirupsen/logrus"
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

func(rfs GrpcRemoteFileServer) Serve(l net.Listener) {
}
