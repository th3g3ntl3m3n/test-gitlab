package main

import (
	"fmt"
	"test-egroup/service"
)

func main() {
	fmt.Println("Hello World!!")
}

func GitlabSvc() string {
	return service.GitlabSvc()
}
