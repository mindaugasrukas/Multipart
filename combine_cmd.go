package main

import (
	"fmt"
	"log"
	"multipart/multipart"
	"os"

	"github.com/spf13/cobra"
)

var combineCmd = &cobra.Command{
	Use:     "combine",
	Aliases: []string{"c"},
	Short:   "",
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		outputFile := os.Stdout
		if output != "" {
			var err error
			outputFile, err = os.Create(output)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			defer outputFile.Close()
		}

		log.Println("Combining files:", args)
		multipart := multipart.New()
		for _, arg := range args {
			err := multipart.AddFile(arg)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		err := multipart.Write(outputFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	combineCmd.Flags().StringP("output", "o", "", "output file")
	rootCmd.AddCommand(combineCmd)
}
