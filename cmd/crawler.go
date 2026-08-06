/*
 * Copyright © 2026 NAME HERE <EMAIL ADDRESS>
 * Generated with `cobra-cli add crawler`
 */
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	crawlerPath string
	crawlerType string
)

// crawlerCmd represents the crawler command
var crawlerCmd = &cobra.Command{
	Use:   "crawler",
	Short: "Run the crawler",
	Long:  `Crawls the given path using the specified crawler type.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Running crawler: path=%s type=%s\n", crawlerPath, crawlerType)
		// TODO: actual crawler logic
		return nil
	},
}

func init() {
	rootCmd.AddCommand(crawlerCmd)

	crawlerCmd.Flags().StringVar(&crawlerPath, "path", "", "path to crawl (required)")
	crawlerCmd.Flags().StringVar(&crawlerType, "type", "", "type of crawler to use (required)")

	crawlerCmd.MarkFlagRequired("path")
	crawlerCmd.MarkFlagRequired("type")
}
