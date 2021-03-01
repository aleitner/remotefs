package main

import (
	"os"

	"remotefs/pkg/server"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "serve",
				Aliases: []string{"s"},
				Usage:   "join a call room",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "directory",
						Value:    "",
						Usage:    "directory path being served",
						Required: true,
					},
					&cli.StringFlag{
						Name:  "address",
						Value: ":8080",
						Usage: "server address",
					},
				},
				Action: func(c *cli.Context) error {
					var logger = log.New()

					rfs, err := server.NewRFS(logger, c.String("directory"), c.String("address"))
					if err != nil {
						return err
					}

					rfs.Serve()

					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}