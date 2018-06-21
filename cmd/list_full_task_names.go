package cmd

import (
	"strings"
	"sort"
	"fmt"
	"bytes"
)

// List full tasks names on one line
func ListFullTaskNames(opts *mykeOpts, tasks []string) error {
	w, err := loadWorkspace(opts.File)
	if err != nil {
		return err
	}

	for _, p := range w.Projects {
		tasks := []string{}
		for t := range p.Tasks {
			if !strings.HasPrefix(t, "_") {
				tasks = append(tasks, t)
			}
		}
		if len(tasks) > 0 {
			sort.Strings(tasks)
			var projectName bytes.Buffer
			projectName.WriteString(p.Name)
			projectName.WriteString("/")
			for _, sortedTask := range tasks {
				var fullTaskName bytes.Buffer
				fullTaskName.Write(projectName.Bytes())
				fullTaskName.WriteString(sortedTask)
				fullTaskName.WriteString(" ")
				fmt.Fprintln(opts.Writer, fullTaskName.String())
			}
		}
	}
	return nil
}
