package gitlab

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/xanzy/go-gitlab"
	"gopkg.in/yaml.v2"
)

// Dockerized path to templates
var projectsYamlPath string = "/go/src/mr_review/config/projects.yaml"

type ActiveProjects struct {
	Projects []Project
}

type Project struct {
	Name string
	Mrs  []*gitlab.MergeRequest
	PID  string
}

func GetOpenMRs(client *gitlab.Client) ActiveProjects {
	projs := getActiveProjs()

	state := "opened"
	scope := "all"
	sort := "asc"

	opts := gitlab.ListProjectMergeRequestsOptions{
		State: &state,
		Scope: &scope,
		Sort:  &sort,
	}

	for i, project := range projs.Projects {
		mrs, _, err := client.MergeRequests.ListProjectMergeRequests(project.PID, &opts)
		if err != nil {
			log.Printf("Error getting merge requests for project %s: %s", project.Name, err)
		}
		p := &projs.Projects[i]
		p.Mrs = mrs
	}

	return projs
}

func getGitlabToken() string {
	t := os.Getenv("GITLAB_TOKEN")
	if t == "" {
		log.Printf("No Gitlab token found - add token to env variable GITLAB_TOKEN")
	}
	return t
}

func NewGitlabClient(baseURL string) *gitlab.Client {
	client := gitlab.NewClient(nil, getGitlabToken())
	if baseURL != "" {
		client.SetBaseURL(baseURL)
	}

	return client
}

func getActiveProjs() ActiveProjects {
	type yamlProj struct {
		Name string
		PID  string
	}

	var y []yamlProj

	source, err := ioutil.ReadFile(projectsYamlPath)
	log.Printf("Source: %s", source)

	err = yaml.Unmarshal([]byte(source), &y)
	if err != nil {
		log.Printf("Error consuming projects YAML file: %s\n", err)
	}

	log.Printf("Projects YAML: %s", y)

	var activeProjs ActiveProjects

	for _, p := range y {
		proj := Project{
			Name: p.Name,
			PID:  p.PID,
		}
		activeProjs.Projects = append(activeProjs.Projects, proj)
	}
	return activeProjs
}
