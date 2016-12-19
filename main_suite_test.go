package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/urfave/cli.v1"
	"github.com/apex/log"
	logcli "github.com/apex/log/handlers/cli"
	"github.com/goeuro/myke/cmd"
	"testing"
	"bytes"
)

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "myke CLI")
}

// Initial boilerlate for CLI BDD
// We can remove other tests (where possible)
// And switch to table-based testing

var _ = Describe("myke", func() {

	var stdout *bytes.Buffer
	var app *cli.App

	BeforeEach(func() {
		stdout = bytes.NewBufferString("")
		app = cmd.NewApp(stdout)
		log.SetHandler(&logcli.Handler{Writer: stdout, Padding: 0})
	})

	Describe("./myke.yml", func() {
		It("list", func() {
			app.Run([]string{""})
			Expect(stdout.String(), MatchRegexp("PROJECT\\s*|\\s*TAGS\\s*|\\s*TASKS"))
			Expect(stdout.String(), MatchRegexp("myke\\s*|\\s*|\\s*test, package"))
		})
	})

})
