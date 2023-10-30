package main

import (
	"fmt"
	"log"
	"os"

	"github.com/shtayeb/go-organizer/organizers"
)

func main() {

	path, _ := os.Getwd()

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

	fmt.Printf("Program run !")

}
