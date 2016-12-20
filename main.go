package main

import (
	"github.com/goeuro/myke/cmd"
	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"os"
)

func main() {
	log.SetHandler(&cli.Handler{Writer: os.Stdout, Padding: 0})
	if err := cmd.NewApp().Run(os.Args); err != nil {
		log.WithError(err).Fatal("error")
	}
}
