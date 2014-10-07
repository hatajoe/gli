package milestones

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/tabwriter"
)

type Env struct {
	Endpoint    string
	Version     int
	TokenSecret string
	Body        string
}

type Assignee struct {
	Username string
}

type Milestone struct {
	Iid      int
	State    string
	Title    string
	Due_date string
}

type Milestones []Milestone

func (m *Milestone) ToLine() string {
	return fmt.Sprintf("#%d\t%s\t%s\t%s", m.Iid, m.State, m.Due_date, m.Title)
}

func (ms *Milestones) EchoLines() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 4, '\t', 0)
	for _, m := range *ms {
		fmt.Fprintln(w, m.ToLine())
	}
	w.Flush()
}

func Describe(projectId int, page int, env Env) (milestones Milestones, err error) {

	if env.Body == "" {
		uri := fmt.Sprintf("http://%s/api/v%d/projects/%d/milestones?private_token=%s&page=%d&per_page=100", env.Endpoint, env.Version, projectId, env.TokenSecret, page)
		resp, err := http.Get(uri)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		env.Body = string(body)
	}

	err = json.Unmarshal([]byte(env.Body), &milestones)
	if err != nil {
		return nil, err
	}
	return milestones, err
}
