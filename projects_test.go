package projects

import "os"
import "strconv"
import "testing"

func TestDescribe(t *testing.T) {
    expectedResponseJson := `
[
  {
    "id": 4,
    "description": null,
    "default_branch": "master",
    "public": false,
    "visibility_level": 0,
    "ssh_url_to_repo": "git@example.com:diaspora/diaspora-client.git",
    "http_url_to_repo": "http://example.com/diaspora/diaspora-client.git",
    "web_url": "http://example.com/diaspora/diaspora-client",
    "owner": {
      "id": 3,
      "name": "Diaspora",
      "created_at": "2013-09-30T13: 46: 02Z"
    },
    "name": "Diaspora Client",
    "name_with_namespace": "Diaspora / Diaspora Client",
    "path": "diaspora-client",
    "path_with_namespace": "diaspora/diaspora-client",
    "issues_enabled": true,
    "merge_requests_enabled": true,
    "wall_enabled": false,
    "wiki_enabled": true,
    "snippets_enabled": false,
    "created_at": "2013-09-30T13: 46: 02Z",
    "last_activity_at": "2013-09-30T13: 46: 02Z",
    "namespace": {
      "created_at": "2013-09-30T13: 46: 02Z",
      "description": "",
      "id": 3,
      "name": "Diaspora",
      "owner_id": 1,
      "path": "diaspora",
      "updated_at": "2013-09-30T13: 46: 02Z"
    },
    "archived": false
  },
  {
    "id": 6,
    "description": null,
    "default_branch": "master",
    "public": false,
    "visibility_level": 0,
    "ssh_url_to_repo": "git@example.com:brightbox/puppet.git",
    "http_url_to_repo": "http://example.com/brightbox/puppet.git",
    "web_url": "http://example.com/brightbox/puppet",
    "owner": {
      "id": 4,
      "name": "Brightbox",
      "created_at": "2013-09-30T13:46:02Z"
    },
    "name": "Puppet",
    "name_with_namespace": "Brightbox / Puppet",
    "path": "puppet",
    "path_with_namespace": "brightbox/puppet",
    "issues_enabled": true,
    "merge_requests_enabled": true,
    "wall_enabled": false,
    "wiki_enabled": true,
    "snippets_enabled": false,
    "created_at": "2013-09-30T13:46:02Z",
    "last_activity_at": "2013-09-30T13:46:02Z",
    "namespace": {
      "created_at": "2013-09-30T13:46:02Z",
      "description": "",
      "id": 4,
      "name": "Brightbox",
      "owner_id": 1,
      "path": "brightbox",
      "updated_at": "2013-09-30T13:46:02Z"
    },
    "archived": false
  }
]
    `
    expectedDescribedStructs := Projects {
        {
            Id: 4, 
            Ssh_url_to_repo: "git@example.com:diaspora/diaspora-client.git", 
            Http_url_to_repo: "http://example.com/diaspora/diaspora-client.git",
        },
        {
            Id: 6, 
            Ssh_url_to_repo: "git@example.com:brightbox/puppet.git", 
            Http_url_to_repo: "http://example.com/brightbox/puppet.git",
        },
    }
    expectedCount := 2

    testedCount := 0
    version, _ := strconv.Atoi(os.Getenv("GITLAB_API_VERSION"))
    env := Env { 
        Endpoint: os.Getenv("GITLAB_API_DOMAIN"),
        Version: version,
        TokenSecret: os.Getenv("GITLAB_API_TOKEN"),
        Body: expectedResponseJson,
    }
    describedProjects, err := Describe(env)
    if err != nil {
        t.Error(err)
    }
    for _, expected := range expectedDescribedStructs {
        for _, described := range describedProjects {
            if (expected.Id == described.Id &&
                expected.Ssh_url_to_repo == described.Ssh_url_to_repo &&
                expected.Http_url_to_repo == described.Http_url_to_repo) {
                testedCount = testedCount + 1
            }
        }
    }

    if expectedCount != testedCount {
        t.Errorf("return object has invalid format")
    }
 
    env = Env { 
        Endpoint: os.Getenv("GITLAB_API_DOMAIN"),
        Version: version,
        TokenSecret: os.Getenv("GITLAB_API_TOKEN"),
    }
    describedProjects, err = Describe(env)
    if err != nil {
        t.Error(err)
    }
}

