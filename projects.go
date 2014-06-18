package projects

type Project struct {
    id int
    ssh_url_to_repo string
    http_url_to_repo string
}

type Projects []Project

func Describe(...string) Projects {

    projects := Projects {}
    return projects
}
