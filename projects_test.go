package projects

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
            id: 4, 
            ssh_url_to_repo: "git@example.com:diaspora/diaspora-client.git", 
            http_url_to_repo: "http://example.com/diaspora/diaspora-client.git",
        },
        {
            id: 6, 
            ssh_url_to_repo: "git@example.com:brightbox/puppet.git", 
            http_url_to_repo: "http://example.com/brightbox/puppet.git",
        },
    }
    expectedCount := 2

    testedCount := 0 
    describedProjects := Describe(expectedResponseJson)
    for _, expected := range expectedDescribedStructs {
        for _, described := range describedProjects {
            if (expected.id == described.id &&
                expected.ssh_url_to_repo == described.ssh_url_to_repo &&
                expected.http_url_to_repo == described.http_url_to_repo) {
                testedCount = testedCount + 1
            }
        }
    }

    if expectedCount != testedCount {
        t.Errorf("return object has invalid format")
    }
}
