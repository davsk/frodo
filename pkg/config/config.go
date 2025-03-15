// 'frodo/pkg/config/config.go'

// Package config contains the init, edit, save, and load of the frodo config.
package config

import (
	"fmt"
	"os/exec"
)

// type Profile contains values of user and run count.
type Profile struct {
	countRuns int
	UserName, UserEmail string
}

// New
func Load() *Profile {
	fmt.Printf("test.")

	// golang check
	if _, err := exec.LookPath("go"); err == nil {
		fmt.Println("Go is already installed.")
		return nil
	} else {
		// install.Golang
	}

	// gh check
	if _, err := exec.LookPath("gh"); err == nil {
		fmt.Println("gh is already installed.")
		return nil
	} else {
		// install.Golang
	}
	return new(Profile)
}
