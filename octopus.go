package main

import (
	"fmt"
	"os/exec"
	"log"
	"flag"
)

func main() {
	pattern, repository := InitFlags()
	flag.Parse()

	lsRemoteCmd := exec.Command("git", "ls-remote", *repository, *pattern)
	branchList, err := lsRemoteCmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", branchList)
}

func InitFlags() (*string, *string) {
	pattern := flag.String("pattern", "", "Branch naming pattern")
	repository := flag.String("repository", ".", "Repository name")
	return pattern, repository
}