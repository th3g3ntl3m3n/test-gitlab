package external

import (
	"context"
	"test-egroup/entity"

	"github.com/machinebox/graphql"
)

type GitlabClient struct {
	client *graphql.Client
}

func InitGitlabClient(url string) GitlabClient {
	return GitlabClient{client: graphql.NewClient(url)}
}

func (glab GitlabClient) LastProjectsGqlAPI(param int) (entity.GitlabGqlResp, error) {
	req := graphql.NewRequest(`
		query last_projects($n: Int = 5) {
			projects(last:$n) {
				nodes {
					name
					description
					forksCount
				}
			}
		}
	`)

	req.Var("n", param)

	ctx := context.Background()

	var resp entity.GitlabGqlResp

	err := glab.client.Run(ctx, req, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
