package service

import (
	"strings"
	"testing"

	"test-egroup/utils"
)

type TestStruct struct {
	Got   int
	Want  int
	Param int
	Name  string
}

func TestGitlabSvc(t *testing.T) {
	var table = []TestStruct{
		{5, 5, 5, "Param with 5"},
		{0, 0, 0, "Param with 0"},
		{0, 0, -1, "Param with less then 0"},
	}

	for _, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			got := GitlabSvc(v.Param)
			var want = 0
			if got.Names != "" {
				want = len(strings.Split(got.Names, utils.DELIM))
			}
			if v.Want != want {
				t.Fatal("Test Failed : Gitlab Svc", v.Param, want)
			}
		})
	}
}
