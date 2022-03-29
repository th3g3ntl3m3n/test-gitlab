package main

import (
	"flag"
	"test-egroup/service"
)

func main() {
	var paramFlat = flag.Int("param", 0, "parameter for last project api")
	flag.Parse()

	service.GitlabSvc(*paramFlat).Print()
}
