package main // import "moul.io/assh/v2/contrib/webapp"

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "bind-address",
			Value: ":8080",
		},
	}
	app.Action = server
	if err := app.Run(os.Args); err != nil {
		log.Fatalf("cannot run app: %v", err)
	}
}

func server(c *cli.Context) error { _ = "STUB: not implemented"; return nil }

// "assh_config": json.AsshConfig,
