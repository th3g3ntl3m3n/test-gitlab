package service

import (
	"test-egroup/entity"
	"test-egroup/repository"
)

func GitlabSvc() entity.SvcResp {
	var response = repository.GitlabRepo()
	var svcResp entity.SvcResp
	for _, v := range response.Nodes {
		svcResp.Names = svcResp.Names + v.Name + ", "
		svcResp.TotalFork += v.ForksCount
	}
	return svcResp
}
