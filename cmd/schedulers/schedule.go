package schedulers

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/google/uuid"
)

var taskMarker = "OrganizerScheduledTask"

func GetExecutablePath() (string, error) {
	executable, err := os.Executable()
	if err != nil {
		return "", err
	}

	// If the executable is a symbolic link, resolve it to get the actual path
	if link, err := os.Readlink(executable); err == nil {
		return filepath.Clean(link), nil
	}

	return filepath.Clean(executable), nil
}

// Schedule a command
func ScheduleCommand(path string, scheduleType string) {
	executablePath, err := GetExecutablePath()
	if err != nil {
		executablePath = "Organizer"
		log.Printf("Error getting executable path defaulting to 'Organizer' %v\n", err)

	}

	switch runtime.GOOS {
	case "windows":
		command := "\"" + executablePath + "\"" + " --path=" + path
		scheduleWindowsTask(command, scheduleType)
	case "linux", "darwin":
		command := "\"" + executablePath + "\"" + " --path=" + path
		scheduleUnixCommand(command, scheduleType)

	default:
		log.Println("Unsupported operating system")

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
		log.Printf("Error scheduling task on Windows: %v\n", err)
	} else {
		log.Printf("Task scheduled on Windows (%s)\n", scheduleType)
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
		log.Println("Invalid schedule type")
		return
	}

	// Include a comment as a marker in the command
	// commentMarker := "# Task Marker: test"
	// cmd.Args = append(cmd.Args, fmt.Sprintf("echo \"%s\" 2>&1 && %s", commentMarker, command))

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Printf("Error scheduling command on Linux/macOS: %v\n", err)
	} else {
		log.Printf("Command scheduled on Linux/macOS (%s)\n", scheduleType)
	}
}

// List scheduled commands
func GetScheduledTasks() string {
	tasks := "Tasks not found"
	switch runtime.GOOS {
	case "windows":
		tasks = listWindowsTasks()
	case "linux", "darwin":
		tasks = listUnixTasks()
	default:
		log.Println("Unsupported operating system")
	}
	return tasks
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
		log.Println("TaskName not found in the CSV headers.")
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

func listWindowsTasks() string {
	// Use schtasks to list tasks on Windows
	// /TN "OrganizerScheduledTask"
	cmd := exec.Command("schtasks", "/query", "/fo", "csv", "/v")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("Error listing tasks on Windows: %v\n", err)
	}

	// fmt.Printf("Scheduled Tasks on Windows:\n%s\n", output)
	// printTasksByMarker(string(output))
	return string(output)
}

func listUnixTasks() string {
	// Use atq to list tasks on Linux/macOS
	cmd := exec.Command("atq")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("Error listing tasks on Linux/macOS: %v\n", err)
	}

	return fmt.Sprintf("Scheduled Tasks on Linux/macOS:\n%s\n", output)
}
