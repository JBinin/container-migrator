package main

import (
	"github.com/JBinin/container-migrator/client"
	"github.com/JBinin/container-migrator/server"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	// todo check criu
	// todo check runc
}

func main() {

	var migratedContainerDir string
	var destination string
	var containerId string
	var othersPath string
	app := &cli.App{
		Name:  "migrate",
		Usage: "Migrate containers.",
		Commands: []*cli.Command{
			{
				Name:  "client",
				Usage: "Running as a client.",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "container_id",
						Usage:       "The container id of migrated container.",
						Destination: &containerId,
						Required:    true,
					},
					&cli.StringFlag{
						Name:        "destination",
						Usage:       "The destination node of migration.",
						Destination: &destination,
						Required:    true,
					},
					&cli.StringFlag{
						Name:        "others_path",
						Usage:       "The file path of config.json and rootfs.",
						Destination: &othersPath,
						Required:    true,
					},
				},
				Subcommands: []*cli.Command{
					{
						Name:  "pre_copy",
						Usage: "Using pre_copy mode.",
						Action: func(context *cli.Context) error {
							// todo
							return client.PreCopy(containerId, destination, othersPath)
						},
					},
					{
						Name:  "post_copy",
						Usage: "Using post_copy mode.",
						Action: func(context *cli.Context) error {
							// todo
							return client.PostCopy(containerId, destination)
						},
					},
				},
			},
			{
				Name:  "server",
				Usage: "Running as a server.",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "migrated_container_dir",
						Usage:       "The directory for saving the migrated container files.",
						Destination: &migratedContainerDir,
					},
				},
				Action: func(context *cli.Context) error {
					server.ListenAndServe(migratedContainerDir)
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
