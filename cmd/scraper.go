/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// scraperCmd represents the scraper command
var scraperCmd = &cobra.Command{
	Use:   "scraper <command>",
	Short: "A tool to scraper and filter out data from html page",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("scraper called \n")
	},
}

func init() {
	rootCmd.AddCommand(scraperCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scraperCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scraperCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
