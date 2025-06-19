package cmd

import (
	"context"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/charmbracelet/fang"
	"github.com/shtayeb/organizer/cmd/organizers"
	"github.com/shtayeb/organizer/cmd/schedulers"
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
	Use:   "organizer",
	Short: "A CLI app to organize your files into folders according to their extensions.",
	Run: func(cmd *cobra.Command, args []string) {
		if path == "" {

			if !workingDirectory {
				log.Panic("Must provide a path or working working directory (-d) flag")
			}

			wdPath, _ := os.Getwd()
			path = wdPath
		}

		log.Printf("Direcotry to be organized: %s \n", path)

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
		// The below could be better

		if weekly || monthly {

			if weekly {
				schedulers.ScheduleCommand(path, "--weekly")
			}
			if monthly {
				schedulers.ScheduleCommand(path, "--monthly")
			}

		}

		log.Printf("Organizer Finished Execution successfully ! \n")
	},
}

var listScheduledCmd = &cobra.Command{
	Use:   "list",
	Short: "List all Organizer scheduled tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		schedulers.ListScheduledTasks()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// Setup logging to file
	executablePath, err := schedulers.GetExecutablePath()
	if err != nil {
		log.Panicf("Error getting executable path and that is BAD: %s", err)
	}

	executableDir := filepath.Dir(executablePath)
	logFilePath := filepath.Join(executableDir, "organizer-cli.log")

	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Println(err.Error())
	}

	defer logFile.Close()

	// redirect all the output to file
	wrt := io.MultiWriter(os.Stdout, logFile)

	// set log out put
	log.SetOutput(wrt)

	// optional: log date-time, filename, and line number
	log.SetPrefix("ORGANIZER CLI: ")

	// Execute the command
	if err := fang.Execute(context.TODO(), rootCmd); err != nil {
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
		BoolVarP(&workingDirectory, "working-dir", "d", false, "Organize working directory")

	rootCmd.Flags().
		BoolVarP(&weekly, "weekly", "w", false, "Schedule the command weekly")

	rootCmd.Flags().
		BoolVarP(&monthly, "monthly", "m", false, "Schedule the command monthly")

	rootCmd.AddCommand(listScheduledCmd)
}
