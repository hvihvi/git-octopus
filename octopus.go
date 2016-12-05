package main

import (
	"fmt"
	"github.com/libgit2/git2go"
	"os"
)

func main() {
	repo, err := git.OpenRepository(".")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if repo.IsBare() {
		fmt.Println("Yep, it's bare.")
	} else {
		fmt.Println("Nope. Not a bare repo.")
	}
}
