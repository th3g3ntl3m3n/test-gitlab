package main

import (
	"test-egroup/entity"
	"test-egroup/service"
	"testing"
)

func TestGitlabSvc(t *testing.T) {
	got := service.GitlabSvc()
	want := entity.SvcResp{Names: "", TotalFork: 0}
	if got.Names != want.Names && got.TotalFork != want.TotalFork {
		t.Fatal("Test Failed : Gitlab Svc")
	}
}
