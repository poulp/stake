package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
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