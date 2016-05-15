package main

import (
	"fmt"
	"os"
	"github.com/libgit2/git2go"
	"time"

	"github.com/poulp/stake/core"
)

const VERSION = "0.1"

func stakeError(err error){
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func walkRepo(repo *git.Repository){

	var repoCommits core.StakeRepository

	var authors core.StakeAuthors
	authors = make(core.StakeAuthors)

	walk, err := repo.Walk()
	stakeError(err)
	walk.PushRange("HEAD..HEAD")

	oid := new(git.Oid)
	/* Iterate through commits */
	for {
		err = walk.Next(oid)
		if err != nil {
			if err.(*git.GitError).Code == git.ErrIterOver {
				break
			} else {
				stakeError(err)
			}
		}

		commit, err := repo.LookupCommit(oid)
		stakeError(err)

		// Repository
		repoCommits.Perform(commit)
		// Authors
		authors.Perform(commit)
	}
	fmt.Println(authors)
	fmt.Println(&repoCommits)
}

func main(){
	var stakeInfo core.StakeInfo
	defer stakeInfo.SetInfo(VERSION, time.Now())

	/* Open repository */
	repo, err := git.OpenRepository(".")
	stakeError(err)
	defer repo.Free()

	/* Main */
	walkRepo(repo)
}