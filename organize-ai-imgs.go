package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func OrganizeBingGeneratedImgs(path string, aiFolderName string) int {
	aiImgDirectory := path + aiFolderName

	// Get list of files in the Downloads path
	entries, err := os.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	count := 0
	for _, entry := range entries {

		if entry.IsDir() {
			continue
		}

		fullFileName := entry.Name()

		startsWithUnderscore := strings.HasPrefix(fullFileName, "_")

		if !startsWithUnderscore {
			continue
		}

		fileName := strings.Split(fullFileName, ".")[0][1:]

		if !IsValidUUID(fileName) {
			continue
		}

		// Do the organization
		_, err := os.Stat(aiImgDirectory)

		if os.IsNotExist(err) {
			// create the directory
			err := os.Mkdir(aiImgDirectory, 0755)

			if err != nil {
				// log.Fatal(err)
				continue
			}

		}

		// move the file
		oldPath := filepath.Join(path, entry.Name())
		newPath := filepath.Join(aiImgDirectory, entry.Name())

		// println(oldPath)
		// println(newPath)
		// println("-----=====-----")

		err = os.Rename(oldPath, newPath)

		if err != nil {
			continue
			// log.Fatalf(errNew.Error())
		}

		count++

	}

	return count
}
