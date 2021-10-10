package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var tbbCmd = &cobra.Command{
	Use:   "tbb",
	Short: "The Blockchain Bar CLI",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func main() {
	tbbCmd.AddCommand(versionCmd)
	tbbCmd.AddCommand(balances())
	tbbCmd.AddCommand(tx())
    tbbCmd.Execute()
}

func incorrectUsageErr() error {
	return fmt.Errorf("incorrect usage")
}