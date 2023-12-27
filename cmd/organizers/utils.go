package organizers

import (
	"os"
	"slices"

	"github.com/google/uuid"
)

var (
	IMAGE_EXTS = []string{
		"png",
		"jpg",
		"webp",
		"gif",
		"jpeg",
		"svg",
		"tiff",
		"psd",
		"cr2",
		"crw",
		"arw",
		"bmp",
	}
	DOCUMENT_EXTS   = []string{"pdf", "doc", "docx", "ppt", "pptx", "xls", "xlsx"}
	PROGRAM_EXTS    = []string{"exe", "apk", "deb", "msi", "dmg"}
	TEXTFILE_EXTS   = []string{"txt", "md"}
	COMPRESSED_EXTS = []string{
		"zip",
		"rar",
		"7z",
		"tgz",
		"tar.gz",
		"txz",
		"arc",
		"arj",
		"as",
		"b64",
		"btoa",
		"bz",
		"bz2",
		"cab",
		"cpt",
		"gz",
		"hqx",
		"iso",
		"lha",
		"lzh",
		"mim",
		"mme",
		"pak",
		"pf",
		"rar",
		"rpm",
		"sea",
		"sit",
		"sitx",
		"tbz",
		"tbz2",
		"tgz",
		"uu",
		"uue",
		"z",
		"zip",
		"zipx",
		"zoo",
	}
)

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

func isCompressedFile(ext string) bool {
	return slices.Contains(COMPRESSED_EXTS, ext)
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
