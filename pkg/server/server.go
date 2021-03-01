package server

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type RemoteFileServer struct {
	directory string
	address string
	logger *log.Logger
}

func NewRFS(logger *log.Logger, dir, address string) (*RemoteFileServer, error) {
	fi, err := os.Stat(dir)
	if err != nil || !fi.Mode().IsDir() {
		return nil, fmt.Errorf("%s is not a valid directory", dir)
	}

	return &RemoteFileServer{
		directory: dir,
		address: address,
		logger:    logger,
	}, nil
}

func(rfs RemoteFileServer) Serve() {
	fs := http.FileServer(http.Dir(rfs.directory))

	http.ListenAndServe(rfs.address, fs)
}
