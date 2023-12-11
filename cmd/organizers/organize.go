package organizers

import (
	"path/filepath"
	"strings"
)

func OrganizeFiles(path, fullFileName string) {

	var folderName string

	fullNameArray := strings.Split(fullFileName, ".")
	fileExt := fullNameArray[len(fullNameArray)-1]

	switch {
	case isImage(fileExt):
		fileName := strings.Split(fullFileName, ".")[0][1:]
		startsWithUnderscore := strings.HasPrefix(fullFileName, "_")

		if startsWithUnderscore && IsValidUUID(fileName) {
			folderName = "AI Images"
		} else {
			folderName = "Images"
		}
	case isDocument(fileExt):
		folderName = "Documents"
	case isProgram(fileExt):
		folderName = "Programs"
	case isTextFile(fileExt):
		folderName = "Text Files"
	case isCompressedFile(fileExt):
		folderName = "Compressed"
	default:
		folderName = "Others"
	}

	directory := path + "/" + folderName
	oldPath := filepath.Join(path, fullFileName)
	newPath := filepath.Join(directory, fullFileName)

	moveFiles(directory, oldPath, newPath)
}
