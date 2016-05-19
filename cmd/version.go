package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
    RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
    Use:   "version",
    Short: "Print the version number of Stake",
    Long:  "Print the version number of Stake",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Stake v0.1")
    },
}