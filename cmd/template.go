package cmd

import (
	"gopkg.in/urfave/cli.v1"
	"github.com/goeuro/myke/core"
	"github.com/apex/log"
	"fmt"
	"io/ioutil"
)

func Template(c *cli.Context) error {
	bytes, err := ioutil.ReadFile(c.String("template"))
	if err != nil {
		log.WithError(err).Fatal("error rendering template")
	}

	rendered, err := core.RenderTemplate(string(bytes), core.OsEnv(), map[string]string{})
	if err != nil {
		log.WithError(err).Fatal("error rendering template")
	}

	fmt.Fprint(c.App.Writer, rendered)
	return nil
}
