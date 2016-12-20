package examples

import (
	. "github.com/goeuro/myke/examples/util"
	"testing"
)

var tests = []TestTable {
	{`heading`, ``, `PROJECT\s*\|\s*TAGS\s*\|\s*TASKS`},
	{`example`, ``, `example\s*\|\s*\|\s*build`},
	// {`env`, ``, `env\s*\|\s*\|\s*env`},
	{`tags1`, ``, `tags1\s*\|\s*tagA, tagB\s*\|\s*tag`},
	{`tags2`, ``, `tags2\s*\|\s*tagB, tagC\s*\|\s*tag`},
}

func Test(t *testing.T) {
	RunCliTests(t, "examples", tests)
}
