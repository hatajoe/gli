package projects

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

type Project struct {
	Id                  int
	Path_with_namespace string
	Web_url             string
	Ssh_url_to_repo     string
}

type Projects []Project

func (p *Project) ToLine() string {
	return fmt.Sprintf("#%d\t%s\t%s\t%s", p.Id, p.Path_with_namespace, p.Web_url, p.Ssh_url_to_repo)
}

func (ps *Projects) EchoLines() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 4, '\t', 0)
	for _, p := range *ps {
		fmt.Fprintln(w, p.ToLine())
	}
	w.Flush()
}

func Describe(page int, env Env) (projects Projects, err error) {

	if env.Body == "" {
		uri := fmt.Sprintf("http://%s/api/v%d/projects?private_token=%s&page=%d", env.Endpoint, env.Version, env.TokenSecret, page)
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
	err = json.Unmarshal([]byte(env.Body), &projects)
	if err != nil {
		return nil, err
	}
	return projects, err
}
