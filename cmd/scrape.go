package cmd

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var zennURL = "https://zenn.dev/"
var maxArticles int

var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "scrape a website and output the data to the terminal",
	Long:  "scrape a website and output the data to the terminal",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Scrape the website
		resp, err := http.Get(zennURL)
		if err != nil {
			return fmt.Errorf("failed to scrape the website: %s", err)
		}
		defer resp.Body.Close()

		// Check if the response status code is not 200 OK
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("failed to scrape the website: %s", resp.Status)
		}

		// Load the HTML document
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to load the HTML document: %s", err)
		}

		// Find and print the article titles and URLs
		doc.Find(".ArticleList_link__4Igs4").Each(func(i int, s *goquery.Selection) {
			if i >= maxArticles {
				return
			}
			title := s.Find("h2").Text()
			articleUrl, _ := s.Attr("href")
			articleUrl, _ = url.JoinPath(zennURL, articleUrl)
			fmt.Printf("・%s\n(%s)\n\n", title, color.HiGreenString(articleUrl))
		})

		return nil
	},
}

func init() {
	rootCmd.AddCommand(scrapeCmd)

	scrapeCmd.Flags().IntVarP(&maxArticles, "max-articles", "m", 10, "Maximum number of articles to scrape")
}
