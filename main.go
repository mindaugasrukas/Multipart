package main

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "multipart",
	Short: "Combine multiple files into a single MIME multipart file",
}

func main() {
	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	// Configure default logger
	// TODO: make this configurable
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
}
