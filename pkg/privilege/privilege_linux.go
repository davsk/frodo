//go:build linux
// +build linux

// file: 'frodo/pkg/privilege/privilege_linux.go'

// by David Lynn Skinner
// on January 10, 2025
// for Davsk Ltd Co

// copyright (c) 2025 Davsk Ltd Co

// package privilege checks for root, sudo, or admin privilege
// and restarts with the requested privilege

package privilege

import (
	"fmt"
	"os"
	"os/exec"
	"log"
	"strings"
)

func Doit(runCommand string) {
	aCommand := strings.Split(runCommand, " ")
	cmd := exec.Command("sudo", aCommand[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}

func becomeAdmin() {
	//   	exe, _ := os.Executable()
	//   	cwd, _ := os.Getwd()
	//   	args := strings.Join(os.Args[1:], " ")
	cmd := exec.Command("sudo", "ls", "-l", "/root")

	// Set the appropriate environment variables
	cmd.Env = append(os.Environ(), "SUDO_ASKPASS=/usr/bin/ssh-askpass")

	// If the command requires input, you can provide it like this:
	// cmd.Stdin = strings.NewReader("your_password\n")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(string(output))
}
