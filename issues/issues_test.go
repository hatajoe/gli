package issues

import "fmt"
import "os"
import "strconv"
import "strings"
import "testing"

func TestDescribe(t *testing.T) {
	expectedResponseJSON := `
[
  {
    "id": 43,
    "iid": 3,
    "project_id": 8,
    "title": "4xx/5xx pages",
    "description": "",
    "labels": [],
    "milestone": null,
    "assignee": null,
    "author": {
      "id": 1,
      "username": "john_smith",
      "email": "john@example.com",
      "name": "john smith",
      "state": "active",
      "created_at": "2012-05-23t08:00:58z"
    },
    "state": "closed",
    "updated_at": "2012-07-02t17:53:12z",
    "created_at": "2012-07-02t17:53:12z"
  },
  {
    "id": 42,
    "iid": 4,
    "project_id": 8,
    "title": "Add user settings",
    "description": "",
    "labels": [
      "feature"
    ],
    "milestone": {
      "id": 1,
      "title": "v1.0",
      "description": "",
      "due_date": "2012-07-20",
      "state": "reopenend",
      "updated_at": "2012-07-04T13:42:48Z",
      "created_at": "2012-07-04T13:42:48Z"
    },
    "assignee": {
      "id": 2,
      "username": "jack_smith",
      "email": "jack@example.com",
      "name": "Jack Smith",
      "state": "active",
      "created_at": "2012-05-23T08:01:01Z"
    },
    "author": {
      "id": 1,
      "username": "john_smith",
      "email": "john@example.com",
      "name": "John Smith",
      "state": "active",
      "created_at": "2012-05-23T08:00:58Z"
    },
    "state": "opened",
    "updated_at": "2012-07-12T13:43:19Z",
    "created_at": "2012-06-28T12:58:06Z"
  }
]
    `
	expectedDescribedStructs := Issues{
		{
			Iid:        3,
			Title:      "4xx/5xx pages",
			Labels:     []string{},
			State:      "closed",
			Created_at: "2012-07-02t17:53:12z",
		},
		{
			Iid:        4,
			Title:      "Add user settings",
			Labels:     []string{"feature"},
			State:      "opened",
			Created_at: "2012-06-28T12:58:06Z",
		},
	}
	expectedCount := 2

	testedCount := 0
	version, _ := strconv.Atoi(os.Getenv("GITLAB_API_VERSION"))
	env := Env{
		Endpoint:    os.Getenv("GITLAB_API_DOMAIN"),
		Version:     version,
		TokenSecret: os.Getenv("GITLAB_API_TOKEN"),
		Body:        expectedResponseJSON,
	}
	describedIssues, err := Describe(8, 0, env)
	if err != nil {
		t.Error(err)
	}
	for _, expected := range expectedDescribedStructs {
		for _, described := range describedIssues {
			if expected.Iid == described.Iid &&
				expected.Title == described.Title &&
				expected.State == described.State &&
				len(expected.Labels) == len(described.Labels) &&
				expected.Created_at == described.Created_at {

				if len(expected.Labels) > 0 {
					for i, e := range expected.Labels {
						if e != described.Labels[i] {
							t.Error("invalid Labels")
						}
					}
				}
				testedCount = testedCount + 1
			}
		}
	}
	if expectedCount != testedCount {
		t.Errorf("return object has invalid format")
	}

	testedCount = 0
	for _, expected := range expectedDescribedStructs {
		for _, described := range describedIssues {
			if expected.Iid == described.Iid {
				e := fmt.Sprintf("#%d\t%s\t%s\t%s\t%s", expected.Iid, expected.Title, strings.Join(expected.Labels, ","), expected.State, expected.Created_at)
				if e == described.ToLine() {
					testedCount = testedCount + 1
				}
			}
		}
	}
	if expectedCount != testedCount {
		t.Errorf("return object has invalid line format")
	}

	env = Env{
		Endpoint:    os.Getenv("GITLAB_API_DOMAIN"),
		Version:     version,
		TokenSecret: os.Getenv("GITLAB_API_TOKEN"),
	}
	describedIssues, err = Describe(1, 0, env)
	if err != nil {
		t.Error(err)
	}
	describedIssues.EchoLines()
}
