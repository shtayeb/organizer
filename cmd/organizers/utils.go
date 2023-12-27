package organizers

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"slices"
	"strings"

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

// Schedule Command

var taskMarker = "OrganizerScheduledTask"

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
	taskName := taskMarker + "-" + uuid.NewString()
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

	// Include a comment as a marker in the command
	// commentMarker := "# Task Marker: test"
	// cmd.Args = append(cmd.Args, fmt.Sprintf("echo \"%s\" 2>&1 && %s", commentMarker, command))


	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error scheduling command on Linux/macOS: %v\n", err)
	} else {
		fmt.Printf("Command scheduled on Linux/macOS (%s)\n", scheduleType)
	}
}

// List scheduled commands
func ListScheduledTasks() {
	switch runtime.GOOS {
	case "windows":
		listWindowsTasks()
	case "linux", "darwin":
		listUnixTasks()
	default:
		fmt.Println("Unsupported operating system")
	}
}

func printTasksByMarker(tasksOutput string) {
	lines := strings.Split(tasksOutput, "\n")
	header := strings.Split(lines[0], ",")
	taskNameIndex := -1

	for i, h := range header {
		if strings.TrimSpace(h) == "\"TaskName\"" {
			taskNameIndex = i
			break
		}
	}

	if taskNameIndex == -1 {
		fmt.Println("TaskName not found in the CSV headers.")
		return
	}

	for _, line := range lines[2:] {
		fields := strings.Split(line, ",")
		if len(fields) > taskNameIndex {
			taskName := strings.Trim(fields[taskNameIndex], "\"")
			if strings.Contains(taskName, taskMarker) {
				fmt.Println("Scheduled Task Information:")
				for i, h := range header {
					fmt.Printf("%s: %s\n", h, strings.Trim(fields[i], "\""))
				}
				fmt.Println(strings.Repeat("-", 30))
			}
		}
	}
}

func listWindowsTasks() {
	// Use schtasks to list tasks on Windows
	// /TN "OrganizerScheduledTask"
	cmd := exec.Command("schtasks", "/query", "/fo", "csv", "/v")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error listing tasks on Windows: %v\n", err)
		return
	}

	// fmt.Printf("Scheduled Tasks on Windows:\n%s\n", output)
	printTasksByMarker(string(output))
}

func listUnixTasks() {
	// Use atq to list tasks on Linux/macOS
	cmd := exec.Command("atq")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error listing tasks on Linux/macOS: %v\n", err)
		return
	}

	fmt.Printf("Scheduled Tasks on Linux/macOS:\n%s\n", output)
	// printTasksByMarker(string(output))
}
