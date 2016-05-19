package cmd

import (
	"time"
	"fmt"
	"os"

	"github.com/libgit2/git2go"
	"github.com/spf13/cobra"

	"github.com/poulp/stake/core"

)

func init() {
	RootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run Stake",
	Long:  "run Stake",
	Run: run,
}

func run(cmd *cobra.Command, args []string){

	if len(args) <= 0 {
		fmt.Println("Expected arg")
		os.Exit(-1)
	}

	if len(args) > 1{
		fmt.Println("Run command accept only one arg")
		os.Exit(-1)
	}

	output := core.GetStakeOutput(args[0])

	var stakeInfo core.StakeInfo
	defer stakeInfo.SetInfo("0.1", time.Now())

	/* Open repository */
	repo, err := git.OpenRepository(".")
	core.StakeError(err)
	defer repo.Free()

	/* Main */
	core.WalkRepo(repo, output)
}