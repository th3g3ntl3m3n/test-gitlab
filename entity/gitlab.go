package entity

import "fmt"

type GitlabGqlResp struct {
	Projects Projects
}

type Projects struct {
	Nodes []Project
}

type Project struct {
	Name        string
	Description string
	ForksCount  int64
}

type SvcResp struct {
	Names     string
	TotalFork int64
}

func (svc SvcResp) Print() {
	fmt.Printf("Names: %s\nTotal Fork: %d\n", svc.Names, svc.TotalFork)
}
