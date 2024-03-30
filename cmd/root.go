package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var filename string

var rootCmd = &cobra.Command{
	Use:   "goscraper",
	Short: "'goscraper' is a CLI tool to scrape a website and save the data to a CSV file.",
	Long:  `'goscraper' is a CLI tool to scrape a website and save the data to a CSV file.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&filename, "filename", "goscrape.csv", "Filename to save the scraped data")
}
