package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/shtayeb/go-organizer/cmd/organizers"
	"github.com/spf13/cobra"
)

var (
	path             string
	weekly           bool
	monthly          bool
	workingDirectory bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "organizer",
	Version: "0.2.3",
	Short:   "A CLI app to organize your files into folders according to their extensions.",
	Run: func(cmd *cobra.Command, args []string) {
		if path == "" {

			if !workingDirectory {
				log.Panic("Must provide a path or working working directory (-d) flag")
			}

			wdPath, _ := os.Getwd()
			path = wdPath
		}

		fmt.Printf("Direcotry: %s \n", path)

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

		// Schedule the command here
		command := "Organizer --path=" + path
		if weekly {
			organizers.ScheduleCommand(command, "--weekly")
		}
		if monthly {
			organizers.ScheduleCommand(command, "--monthly")
		}

		fmt.Printf("Organizer Finished Execution ! \n")
	},
}

var listScheduledCmd = &cobra.Command{
	Use:   "list",
	Short: "List all Organizer scheduled tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		organizers.ListScheduledTasks()
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

func init() {
	rootCmd.Flags().
		StringVarP(
			&path,
			"path",
			"p",
			"",
			"Absolute path to the directory you want to organize. Default is working directory.",
		)

	rootCmd.Flags().
		BoolVarP(&workingDirectory, "woking-dir", "d", false, "Organize working directory")

	rootCmd.Flags().
		BoolVarP(&weekly, "weekly", "w", false, "Schedule the command weekly")

	rootCmd.Flags().
		BoolVarP(&monthly, "monthly", "m", false, "Schedule the command monthly")

	rootCmd.AddCommand(listScheduledCmd)
}
