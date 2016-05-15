package core

import (
	"github.com/libgit2/git2go"
	"fmt"
)

type StakeAuthor struct {
	Name string
	Mail string
	Total_commits int
}

func (sa *StakeAuthor) String()string{
	return fmt.Sprintf("Name : %s, Mail : %s, Commits : %d", sa.Name, sa.Mail, sa.Total_commits)
}

type StakeAuthors map[string]*StakeAuthor

func (sa StakeAuthors) AddAuthorFromSignature(signature *git.Signature){
	author, exist := sa[signature.Email]
	if exist {
		author.Total_commits++
	} else {
		var new_author = &StakeAuthor{signature.Name, signature.Email, 1}
		sa[signature.Email] = new_author
	}
}

func (sa StakeAuthors) Perform(commit *git.Commit){
	sa.AddAuthorFromSignature(commit.Author())
}