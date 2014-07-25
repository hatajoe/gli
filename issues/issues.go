package issues

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

type Issue struct {
	Iid        int
	State      string
	Created_at string
	Title      string
}

type Issues []Issue

func (i *Issue) ToLine() string {
	return fmt.Sprintf("#%d\t%s\t%s\t%s", i.Iid, i.State, i.Created_at, i.Title)
}

func (is *Issues) EchoLines() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 4, '\t', 0)
	for _, i := range *is {
		fmt.Fprintln(w, i.ToLine())
	}
	w.Flush()
}

func Describe(projectId int, page int, env Env) (issues Issues, err error) {

	if env.Body == "" {
		uri := fmt.Sprintf("http://%s/api/v%d/projects/%d/issues?private_token=%s&page=%d&per_page=100", env.Endpoint, env.Version, projectId, env.TokenSecret, page)
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

	err = json.Unmarshal([]byte(env.Body), &issues)
	if err != nil {
		return nil, err
	}
	return issues, err
}
