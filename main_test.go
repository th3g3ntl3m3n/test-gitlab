package main

import "testing"

func TestGitlabSvc(t *testing.T) {
	if GitlabSvc() != "Gitlab Svc" {
		t.Fatal("Test Failed : Gitlab Svc")
	}
}
