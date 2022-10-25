package main

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"io"
	"os"
)

func main() {
	stat, err := os.Stat("./")
	repo, err := git.PlainOpen(stat.Name())
	if err != nil {
		panic(err)
	}
	fmt.Println("## COMMITS")
	commitIter, err := repo.CommitObjects()
	if err != nil {
		panic(err)
	}
	for cmt, err := commitIter.Next(); err != io.EOF; cmt, err = commitIter.Next() {
		if err != nil {
			panic(err)
		}
		fmt.Println(cmt.Message)
	}

	fmt.Println("\n\n## TAGS")
	tagIter, err := repo.Tags()
	if err != nil {
		panic(err)
	}
	for tag, err := tagIter.Next(); err != io.EOF; tag, err = tagIter.Next() {
		if err != nil {
			panic(err)
		}
		fmt.Printf("%10s : %s\n", tag.Name().Short(), tag.Hash())
	}
}
