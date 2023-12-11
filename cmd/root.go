package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/shtayeb/go-organizer/cmd/organizers"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "organizer",
	Version: "0.1.1",
	Short:   "A CLI app to organize your files into folders according to their extensions.",
	Run: func(cmd *cobra.Command, args []string) {

		path, _ := os.Getwd()

		fmt.Printf("Working Direcotry: %s", path)

		// Get list of files in the working directory
		entries, err := os.ReadDir(path)

		if err != nil {
			log.Fatal(err)
		}

		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}

			fullFileName := entry.Name()
			organizers.OrganizeFiles(path, fullFileName)

		}

		fmt.Printf("Organizer Finished Execution !")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
