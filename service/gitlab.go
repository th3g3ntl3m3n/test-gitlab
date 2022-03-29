package service

import (
	"log"
	"strings"
	"test-egroup/entity"
	"test-egroup/repository"
	"test-egroup/utils"
)

func GitlabSvc(param int) entity.SvcResp {
	if param <= 0 {
		log.Println("Parameter need to be greater than 0")
		return entity.SvcResp{}
	}
	var response = repository.GitlabRepo(param)

	var names []string
	var totalCount int64
	for _, v := range response.Nodes {
		names = append(names, v.Name)
		totalCount += v.ForksCount
	}
	return entity.SvcResp{Names: strings.Join(names, utils.DELIM), TotalFork: totalCount}
}
