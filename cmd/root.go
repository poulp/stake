package cmd

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/libgit2/git2go"

	"github.com/poulp/stake/core"
)

var RootCmd = &cobra.Command{
    Use:   "stake",
    Short: "Stake is a very fast static site generator",
    Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
    Run: rootFunc,
}

func rootFunc(cmd *cobra.Command, args []string){
	var stakeInfo core.StakeInfo
	defer stakeInfo.SetInfo("0.1", time.Now())

	/* Open repository */
	repo, err := git.OpenRepository(".")
	core.StakeError(err)
	defer repo.Free()

	/* Main */
	core.WalkRepo(repo)
}