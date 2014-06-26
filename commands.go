package main

import (
	"github.com/codegangsta/cli"
	"log"
	"os"
)

var Commands = []cli.Command{
	commandIssues,
}

var commandIssues = cli.Command{
	Name:  "issues",
	Usage: "",
	Description: `
`,
	Action: doIssues,
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func doIssues(c *cli.Context) {
}
