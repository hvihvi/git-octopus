package main

import (
	"fmt"
	"os/exec"
	"log"
	"flag"
)

// Command-Line flags.
var (
	pattern = flag.String("pattern", "", "Branch naming pattern")
	repository = flag.String("repository", ".", "Repository name")
)

func main() {
	flag.Parse()

	lsRemoteCmd := exec.Command("git", "ls-remote", *repository, *pattern)
	branchList, err := lsRemoteCmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", branchList)
}
