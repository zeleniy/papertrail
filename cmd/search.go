/*
 * Copyright © 2026 NAME HERE <EMAIL ADDRESS>
 * Generated with `cobra-cli add search`
 */
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	query string
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search the index",
	Long:  `Look up the index by the user query`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("search called")
		fmt.Printf("Running search: query=%s\n", query)
		// TODO: actual search logic
		return nil
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	searchCmd.Flags().StringVar(&query, "query", "", "query string (required)")

	searchCmd.MarkFlagRequired("query")
}
