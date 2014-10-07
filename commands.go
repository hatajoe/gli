package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"os"
	"strconv"

	"github.com/hatajoe/gli/issues"
	"github.com/hatajoe/gli/milestones"
	"github.com/hatajoe/gli/projects"
)

// Commands is available command list
var Commands = []cli.Command{
	commandProjects,
	commandMilestones,
	commandIssues,
}

var commandProjects = cli.Command{
	Name:  "projects",
	Usage: "",
	Description: `
`,
	Action: doProjects,
}

var commandMilestones = cli.Command{
	Name:  "milestones",
	Usage: "",
	Description: `
`,
	Action: doMilestones,
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

func doProjects(c *cli.Context) {
	version, _ := strconv.Atoi(os.Getenv("GITLAB_API_VERSION"))
	env := projects.Env{
		Endpoint:    os.Getenv("GITLAB_API_DOMAIN"),
		Version:     version,
		TokenSecret: os.Getenv("GITLAB_API_TOKEN"),
	}
	describedProjects, err := projects.Describe(1, env)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	page := 2
	for true {
		ps, err := projects.Describe(page, env)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if len(ps) <= 0 {
			break
		}
		for _, p := range ps {
			describedProjects = append(describedProjects, p)
		}
		page++
	}

	describedProjects.EchoLines()
}

func doIssues(c *cli.Context) {
	if len(c.Args()) <= 0 {
		fmt.Println("need to project ID")
		os.Exit(1)
	}
	projectID, err := strconv.Atoi(c.Args()[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	version, _ := strconv.Atoi(os.Getenv("GITLAB_API_VERSION"))
	env := issues.Env{
		Endpoint:    os.Getenv("GITLAB_API_DOMAIN"),
		Version:     version,
		TokenSecret: os.Getenv("GITLAB_API_TOKEN"),
	}
	describedIssues, err := issues.Describe(projectID, 1, env)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	page := 2
	for true {
		is, err := issues.Describe(projectID, page, env)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if len(is) <= 0 {
			break
		}
		for _, i := range is {
			describedIssues = append(describedIssues, i)
		}
		page++
	}

	describedIssues.EchoLines()
}

func doMilestones(c *cli.Context) {
	if len(c.Args()) <= 0 {
		fmt.Println("need to project ID")
		os.Exit(1)
	}
	projectID, err := strconv.Atoi(c.Args()[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	version, _ := strconv.Atoi(os.Getenv("GITLAB_API_VERSION"))
	env := milestones.Env{
		Endpoint:    os.Getenv("GITLAB_API_DOMAIN"),
		Version:     version,
		TokenSecret: os.Getenv("GITLAB_API_TOKEN"),
	}
	describedMilestones, err := milestones.Describe(projectID, 1, env)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	page := 2
	for true {
		is, err := milestones.Describe(projectID, page, env)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if len(is) <= 0 {
			break
		}
		for _, i := range is {
			describedMilestones = append(describedMilestones, i)
		}
		page++
	}

	describedMilestones.EchoLines()
}
