package organizers

import (
	"path/filepath"
	"strings"
)

func OrganizeFiles(path, fullFileName string) {

	folderName := "Others"

	fullNameArray := strings.Split(fullFileName, ".")
	fileExt := fullNameArray[len(fullNameArray)-1]

	if isImage(fileExt) {
		fileName := strings.Split(fullFileName, ".")[0][1:]
		startsWithUnderscore := strings.HasPrefix(fullFileName, "_")

		if startsWithUnderscore && IsValidUUID(fileName) {
			folderName = "AI Images"
		} else {
			folderName = "Images"
		}
	} else if isDocument(fileExt) {
		folderName = "Documents"
	} else if isProgram(fileExt) {
		folderName = "Programs"
	} else if isTextFile(fileExt) {
		folderName = "Text Files"
	}

	directory := path + folderName
	oldPath := filepath.Join(path, fullFileName)
	newPath := filepath.Join(directory, fullFileName)

	moveFiles(directory, oldPath, newPath)
}
