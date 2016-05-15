package core

import (
	"github.com/libgit2/git2go"
	"fmt"
	"os"
)

func StakeError(err error){
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func WalkRepo(repo *git.Repository){

	repoCommits := NewStakeRepository()

	var authors StakeAuthors
	authors = make(StakeAuthors)

	walk, err := repo.Walk()
	StakeError(err)
	walk.PushRange("HEAD..HEAD")

	oid := new(git.Oid)
	/* Iterate through commits */
	for {
		err = walk.Next(oid)
		if err != nil {
			if err.(*git.GitError).Code == git.ErrIterOver {
				break
			} else {
				StakeError(err)
			}
		}

		commit, err := repo.LookupCommit(oid)
		StakeError(err)

		// Repository
		repoCommits.Perform(commit)
		// Authors
		authors.Perform(commit)
	}
	fmt.Println(authors)
	fmt.Println(repoCommits)
}