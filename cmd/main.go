package main

import (
	"fmt"

	gohelloworld "github.com/i-am-david-fernandez/go-hello-world/pkg"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println("Version: ", gohelloworld.VersionGitCommit)
}
