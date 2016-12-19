package main

import (
	"gopkg.in/urfave/cli.v1"
	"github.com/goeuro/myke/cmd"
	"github.com/apex/log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "myke"
	app.Version = cmd.Version()
	app.Usage = "make with yml"
	app.Action = cmd.Action
	app.Flags = []cli.Flag {
		cli.StringFlag{
			 Name: "f, file",
			 Value: "myke.yml",
			 Usage: "`yml-file` to load",
		},
		cli.StringFlag{
			 Name: "template",
			 Usage: "render template `tpl-file` (will not run any command)",
		},
		cli.BoolFlag{
			 Name: "license",
			 Usage: "show license",
		},
		cli.StringFlag{
			Name: "loglevel",
			Value: "info",
			Usage: "log level, one of debug|`info`|warn|error|fatal",
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.WithError(err).Fatal("error")
	}
}
