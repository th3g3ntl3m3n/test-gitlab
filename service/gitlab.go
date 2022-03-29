package service

import (
	"strings"
	"test-egroup/entity"
	"test-egroup/repository"
)

func GitlabSvc() entity.SvcResp {
	var response = repository.GitlabRepo()

	var names []string
	var totalCount int64
	for _, v := range response.Nodes {
		names = append(names, v.Name)
		totalCount += v.ForksCount
	}
	return entity.SvcResp{Names: strings.Join(names, ","), TotalFork: totalCount}
}
