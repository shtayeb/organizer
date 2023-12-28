package schedulers

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/google/uuid"
)

var taskMarker = "OrganizerScheduledTask"

func getExecutablePath() (string, error) {
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
	executablePath, err := getExecutablePath()
	if err != nil {
		executablePath = "Organizer"
		fmt.Printf("Error getting executable path defaulting to 'Organizer' %v\n", err)
	}

	switch runtime.GOOS {
	case "windows":
		command := "\"" + executablePath + "\"" + " --path=" + path
		scheduleWindowsTask(command, scheduleType)
	case "linux", "darwin":
		command := "\"" + executablePath + "\"" + " --path=" + path
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

	// fmt.Printf("Scheduled Tasks on Linux/macOS:\n%s\n", output)
	printTasksByMarker(string(output))
}
