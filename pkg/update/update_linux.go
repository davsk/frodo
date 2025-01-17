// +build linux
// file: 'frodo/pkg/update/update_linux.go'

// by David Lynn Skinner
// on January 10, 2025
// for Davsk Ltd Co

// copyright (c) 2025 Davsk Ltd Co

// package update contains functions to update apps on dev system
package update

import (
	"fmt"
	"os"
	"os/exec"
)

func System() {
	cmd := exec.Command("sudo", "pacman", "-Syyu")
}
