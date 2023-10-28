package organizers

import (
	"os"
	"slices"

	"github.com/google/uuid"
)

var IMAGE_EXTS = []string{"png", "jpg", "webp", "gif", "jpeg", "svg"}
var DOCUMENT_EXTS = []string{"pdf", "doc", "docx", "ppt", "pptx"}
var PROGRAM_EXTS = []string{"exe", "apk", "deb", "msi"}
var TEXTFILE_EXTS = []string{"txt", "md"}

func isImage(ext string) bool {
	return slices.Contains(IMAGE_EXTS, ext)
}

func isDocument(ext string) bool {
	return slices.Contains(DOCUMENT_EXTS, ext)
}

func isProgram(ext string) bool {
	return slices.Contains(PROGRAM_EXTS, ext)
}

func isTextFile(ext string) bool {
	return slices.Contains(TEXTFILE_EXTS, ext)
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func moveFiles(directory, oldPath, newPath string) {
	_, err := os.Stat(directory)

	if os.IsNotExist(err) {
		err := os.Mkdir(directory, 0755)

		if err != nil {
			return
		}

	}

	err = os.Rename(oldPath, newPath)

	if err != nil {
		return
	}
}
