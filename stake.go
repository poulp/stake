package main

import (
	"fmt"
	"os"

	"github.com/poulp/stake/cmd"
)

const VERSION = "0.1"

func main(){
	if err := cmd.RootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(-1)
    }
}