package main

import (
	"fmt"
	"github.com/letieu/fcoin/database"
	"github.com/spf13/cobra"
	"os"
)

const flagFrom string = "from"
const flagTo string = "to"
const flagValue string = "value"
const flagData string = "data"

func tx() *cobra.Command{
	var txsCmd = &cobra.Command{
		Use:   "tx",
		Short: "transaction",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	txsCmd.AddCommand(txAdd())
	return txsCmd
}

func txAdd() *cobra.Command{
	var txAddCmd = &cobra.Command{
		Use:   "add",
		Short: "Add Transaction",
		Run: func(cmd *cobra.Command, args []string) {
			from, _ := cmd.Flags().GetString(flagFrom)
			to, _ := cmd.Flags().GetString(flagTo)
			value, _ := cmd.Flags().GetUint(flagValue)
			data, _ := cmd.Flags().GetString(flagData)

			fromAcc := database.Account(from)
			toAcc := database.Account(to)
			tx := database.Tx{From: fromAcc, To: toAcc, Data: data, Value: value}

			state, err := database.NewStateFromDisk()
			defer state.Close()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			err = state.Add(tx)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			snapshot, err := state.Persist()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			fmt.Printf("New snapshot is %x \n", snapshot)
			fmt.Println("Tx added to the ledger")
		},
	}

	txAddCmd.Flags().String(flagFrom, "", "From what account to send tokens")
	txAddCmd.MarkFlagRequired(flagFrom)

	txAddCmd.Flags().String(flagTo, "", "To what account to send tokens")
	txAddCmd.MarkFlagRequired(flagTo)

	txAddCmd.Flags().Uint(flagValue, 0, "How many tokens to send")
	txAddCmd.MarkFlagRequired(flagValue)

	txAddCmd.Flags().String(flagData, "", "Is reward or not")

	return txAddCmd
}