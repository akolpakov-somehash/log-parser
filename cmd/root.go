/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "log-parser",
	Short: "Log Parser utility",
	Long:  `Log Parser utility.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	var countStatusCodes bool
	var countUniqueIPs bool
	var filterIP string
	var averageBytes bool
	var topUrls bool
	var startDate string
	var endDate string
	var outputFormat string
	var outputFile string

	rootCmd.PersistentFlags().BoolVarP(&countStatusCodes, "count-status-codes", "c", false, "Count status codes")
	rootCmd.PersistentFlags().BoolVarP(&countUniqueIPs, "count-unique-ips", "u", false, "Count unique IPs")
	rootCmd.PersistentFlags().StringVarP(&filterIP, "filter-ip", "f", "", "Filter by IP")
	rootCmd.PersistentFlags().BoolVarP(&averageBytes, "average-bytes", "a", false, "Calculate the average number of bytes sent per request.")
	rootCmd.PersistentFlags().BoolVarP(&topUrls, "top-urls", "t", false, "Top URLs")
	rootCmd.PersistentFlags().StringVarP(&startDate, "start-date", "s", "", "Start date")
	rootCmd.PersistentFlags().StringVarP(&endDate, "end-date", "e", "", "End date")
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "output-format", "o", "json", "Output format")
	rootCmd.PersistentFlags().StringVarP(&outputFile, "output-file", "O", "", "Output file")
}
