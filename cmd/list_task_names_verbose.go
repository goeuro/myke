package cmd

import (
	"strings"
	"sort"
	"fmt"
	"bytes"
)

// List full tasks names on one line
func ListTaskNamesVerbose(opts *mykeOpts, tasks []string) error {
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
			projectTags := strings.Join(p.Tags, ",")
			for _, sortedTask := range tasks {
				var outputLine bytes.Buffer
				outputLine.Write(projectName.Bytes())
				outputLine.WriteString(sortedTask)
				outputLine.WriteString("\t")
				outputLine.WriteString(projectTags)
				outputLine.WriteString("\t")
				task := p.Tasks[sortedTask]
				outputLine.WriteString(task.Desc)
				outputLine.WriteString("\t")
				outputLine.WriteString(p.Src)
				fmt.Fprintln(opts.Writer, outputLine.String())
			}
		}
	}
	return nil
}
