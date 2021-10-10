package main

import (
	"fmt"
	"github.com/letieu/fcoin/database"
	"github.com/spf13/cobra"
	"os"
)

func balances() *cobra.Command {
	var balancesCmd = &cobra.Command{
		Use:   "balances",
		Short: "Interact with balances (list...).",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	balancesCmd.AddCommand(balanceList())
	return balancesCmd
}

func balanceList() *cobra.Command {
	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "List balances",
		Run: func(cmd *cobra.Command, args []string) {
			state, err := database.NewStateFromDisk()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			defer state.Close()

			fmt.Printf("State snapshot: %x \n", state.LatestSnapshot())

			fmt.Println("Accounts balances:")
			fmt.Println("__________________")
			fmt.Println("")
			for account, balance := range state.Balances {
				fmt.Printf(fmt.Sprintf("%s %d \n", account, balance))
			}
		},
	}

	return listCmd
}
