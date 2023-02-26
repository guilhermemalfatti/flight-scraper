/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	googleScraper "github.com/gmalfatti/flight-scraper/googleSccraper"
	"github.com/spf13/cobra"
)

// googleCmd represents the google command
var googleCmd = &cobra.Command{
	Use:   "google",
	Short: "Fetch google flight page and scrape data.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("google called")

		googleScraper.NewScraper().Info()
		/*url := "https://www.google.com/travel/flights/search?tfs=CBwQAhojagwIAhIIL20vMDUycDcSCjIwMjMtMDMtMjFyBwgBEgNQT0EaI2oHCAESA1BPQRIKMjAyMy0wMy0zMXIMCAISCC9tLzA1MnA3cAGCAQsI____________AUABSAGYAQE"
		scraper := googleScraper.NewScraper()
		err := scraper.CreateDocFromURL(url)
		if err != nil {
			return fmt.Errorf("Error %s", err)
		}

		fmt.Printf("todo")*/

		return nil
	},
}

func init() {
	scraperCmd.AddCommand(googleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// googleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// googleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
