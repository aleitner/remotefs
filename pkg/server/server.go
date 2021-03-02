package server

import (
	"fmt"
	"github.com/soheilhy/cmux"
	"net"

	log "github.com/sirupsen/logrus"
)

type RemoteFileServer struct {
	httpServer *HttpRemoteFileServer
	grpcServer *GrpcRemoteFileServer
	logger *log.Logger
	address string
}

func NewRFS(logger *log.Logger, dir, address string) (*RemoteFileServer, error) {
	httpServer, err := NewHttpRFS(logger, dir)
	if err != nil {
		return nil, fmt.Errorf("Failed to create HTTP Server: %s", err)
	}

	grpcServer, err := NewGrpcRFS(logger, dir)
	if err != nil {
		return nil, fmt.Errorf("Failed to create GRPC Server: %s", err)
	}

	return &RemoteFileServer{
		address: address,
		logger:    logger,
		httpServer: httpServer,
		grpcServer: grpcServer,
	}, nil
}

func(rfs RemoteFileServer) Serve() {
	l, err := net.Listen("tcp", rfs.address)
	if err != nil {
		rfs.logger.Error(err)
		return
	}

	// Create a cmux.
	m := cmux.New(l)

	// First grpc, then HTTP, and otherwise Go RPC/TCP.
	grpcL := m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	httpL := m.Match(cmux.HTTP1Fast())

	go rfs.httpServer.Serve(httpL)
	go rfs.grpcServer.Serve(grpcL)

	m.Serve()
}

