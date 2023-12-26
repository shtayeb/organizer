package organizers

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
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

func ScheduleCommand(command string, scheduleType string) {
	switch runtime.GOOS {
	case "windows":
		scheduleWindowsTask(command, scheduleType)
	case "linux", "darwin":
		scheduleUnixCommand(command, scheduleType)
	default:
		fmt.Println("Unsupported operating system")
	}
}

func scheduleWindowsTask(command string, scheduleType string) {
	// Use schtasks to schedule a command on Windows
	taskName := "OrganizerScheduledTask"
	schedule := "/sc"
	interval := "1"

	switch scheduleType {
	case "--weekly":
		interval = "weekly"
	case "--monthly":
		interval = "monthly"
	}

	cmd := exec.Command("schtasks", "/create", "/tn", taskName, "/tr", command, schedule, interval)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error scheduling task on Windows: %v\n", err)
	} else {
		fmt.Printf("Task scheduled on Windows (%s)\n", scheduleType)
	}
}

func scheduleUnixCommand(command string, scheduleType string) {
	// Use at to schedule a command on Linux/macOS
	cmd := exec.Command("at", "now")

	switch scheduleType {
	case "--weekly":
		cmd.Args = append(cmd.Args, "+1 week")
	case "--monthly":
		cmd.Args = append(cmd.Args, "+1 month")
	default:
		fmt.Println("Invalid schedule type")
		return
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Printf("Error creating stdin pipe: %v\n", err)
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Printf("Error starting command: %v\n", err)
		return
	}

	// Write the command to the at command's stdin
	_, err = stdin.Write([]byte(command + "\n"))
	if err != nil {
		fmt.Printf("Error writing to stdin: %v\n", err)
	}

	if err := stdin.Close(); err != nil {
		fmt.Printf("Error closing stdin: %v\n", err)
	}

	if err := cmd.Wait(); err != nil {
		fmt.Printf("Error waiting for command: %v\n", err)
	} else {
		fmt.Printf("Command scheduled on Linux/macOS (%s)\n", scheduleType)
	}
}
