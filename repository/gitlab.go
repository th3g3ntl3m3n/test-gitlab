package repository

import "test-egroup/entity"

func GitlabRepo() entity.Projects {
	return entity.Projects{
		Nodes: []entity.Project{
			{Name: "test", Description: "test", ForksCount: 1},
		},
	}
}
