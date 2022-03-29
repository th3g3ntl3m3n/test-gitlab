package repository

import (
	"log"
	"os"
	"test-egroup/entity"
	"test-egroup/external"
	"test-egroup/utils"
)

func GitlabRepo(param int) *entity.Projects {
	url := os.Getenv(utils.API_URL)
	if url == "" {
		url = "https://gitlab.com/api/graphql"
	}
	resp, err := external.InitGitlabClient(url).LastProjectsGqlAPI(param)
	if err != nil {
		log.Println("Something went wrong : ", err)
		return nil
	}

	return &resp.Projects
}
