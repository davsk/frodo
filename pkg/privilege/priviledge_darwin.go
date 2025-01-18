//go:build darwin
// +build darwin

// file: 'frodo/pkg/privilege/privilege_darwin.go'

// by David Lynn Skinner
// on January 10, 2025
// for Davsk Ltd Co

// copyright (c) 2025 Davsk Ltd Co

// package privilege checks for root, sudo, or admin privilege
// and restarts with the requested privilege

package privilege

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func becomeAdmin() {
	//   	exe, _ := os.Executable()
	//   	cwd, _ := os.Getwd()
	//   	args := strings.Join(os.Args[1:], " ")
}

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
