package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/urfave/cli.v1"
	"github.com/goeuro/myke/cmd"
	"testing"
	"bytes"
	"strings"
)

var testtable = []struct {
	name string
	arg  string
	out  string
}{
	{`heading`, ``, `PROJECT\s*|\s*TAGS\s*|\s*TASKS`},
	{`myke`, ``, `myke\s*|\s*|\s*test, package`},
	{`example`, ``, `example\s*|\s*|\s*build`},
	{`env`, ``, `env\s*|\s*|\s*env`},
	{`tags1`, ``, `tags1\s*|\s*tagA, tagB\s*|\s*tag`},
	{`tags2`, ``, `tags2\s*|\s*tagB, tagC\s*|\s*tag`},
}

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "myke CLI")
}

var _ = Describe("myke", func() {
	var stdout *bytes.Buffer
	var app *cli.App

	BeforeEach(func() {
		stdout = bytes.NewBufferString("")
		app = cmd.NewApp()
		app.Writer = stdout
	})

	for i, _:= range testtable {
		tt := testtable[i]
		It(tt.name, func() {
			app.Run(strings.Split(tt.arg, " "))
			Expect(stdout.String()).To(MatchRegexp(tt.out))
		})
	}
})
