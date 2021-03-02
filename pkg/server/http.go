package server

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"
	"os"
)

type HttpRemoteFileServer struct {
	directory string
	logger *log.Logger
}

func NewHttpRFS(logger *log.Logger, dir string) (*HttpRemoteFileServer, error) {
	fi, err := os.Stat(dir)
	if err != nil || !fi.Mode().IsDir() {
		return nil, fmt.Errorf("%s is not a valid directory", dir)
	}

	return &HttpRemoteFileServer{
		directory: dir,
		logger:    logger,
	}, nil
}

func(rfs HttpRemoteFileServer) Serve(l net.Listener) {
	fs := http.FileServer(http.Dir(rfs.directory))

	http.Serve(l, fs)
}
