package main

import (
	"fmt"
	"log"
	"os"

	"github.com/shtayeb/go-organizer/organizers"
)

func main() {

	path := "C:\\Users\\shahr\\Downloads\\"

	// Get list of files in the Downloads path
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
