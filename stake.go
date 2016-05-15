package main

import (
	"fmt"
	"os"
	"github.com/libgit2/git2go"
)

func stakeError(err error){
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func walkRepo(repo *git.Repository){
	walk, err := repo.Walk()
	stakeError(err)
	walk.PushRange("HEAD..HEAD")
	commits := make([]*git.Commit, 0, 1)

	oid := new(git.Oid)
	/* Iterate throught commits */
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

		commits = append(commits, commit)
	}

	fmt.Println(commits)
}

func main(){
	repo, err := git.OpenRepository(".")
	stakeError(err)
	defer repo.Free()

	walkRepo(repo)
}