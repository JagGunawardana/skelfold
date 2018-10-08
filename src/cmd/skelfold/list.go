package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all templates",
	Long:  `List all template names and descriptions`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Listing:\n- one\n-two\n")
	},
}
