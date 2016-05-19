package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
    Use:   "stake",
    Short: "Stake git statistics tool",
    Long: "Stake git statistics tool",
    Run: rootFunc,
}

func rootFunc(cmd *cobra.Command, args []string) {
	fmt.Println("ok")
}