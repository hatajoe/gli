package projects

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

type Env struct {
    Endpoint    string
    Version     int
    TokenSecret string
    Body        string
}

type Project struct {
    Id int
    Ssh_url_to_repo string
    Http_url_to_repo string
}

type Projects []Project

func Describe(env Env) (projects Projects, err error) {

    if env.Body == "" {
        uri := fmt.Sprintf("http://%s/api/v%d/projects?private_token=%s", env.Endpoint, env.Version, env.TokenSecret)
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
