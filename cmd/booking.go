/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/gmalfatti/flight-scraper/bookingScraper"
	"github.com/spf13/cobra"
)

// bookingCmd represents the booking command
var bookingCmd = &cobra.Command{
	Use:   "booking",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("booking called")

		bookingScraper.NewScraper().Info()
	},
}

func init() {
	scraperCmd.AddCommand(bookingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bookingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bookingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
