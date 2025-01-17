// 'frodo/pkg/install/install.go'

// by David Lynn Skinner
// on January 10, 2025
// for Davsk Ltd Co

// copyright (c) 2025 Davsk Ltd Co

// package install contains functions to install apps on dev system
package install // import "go.davsk.net/frodo/pkg/install"

import (
	"fmt"
	"os"
	"os/exec"
)

// func must receives error and panics
func must(err error) {
	if err != nil {
		fmt.Printf("err: %s\n", err)
		panic(err)
	}
}

// func assert receives a condition, format, and args
// returns a panic if condition is not true
func assert(cond bool, format string, args ...interface{}) {
	if cond {
		return
	}
	s := fmt.Sprintf(format, args...)
	panic(s)
}

// func logf receives a format and args to print
// a centralized place allows us to tweak logging, if need be
func logf(format string, args ...interface{}) {
	if len(args) == 0 {
		fmt.Print(format)
		return
	}
	fmt.Printf(format, args...)
}

// func hashicorp installs the hashicorp  tools
func Hashicorp() {
	cmd := exec.Command("brew", "upgrade")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	must(err)
	cmd.Wait()
}
/*
package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// checkAndInstall checks if a command exists and installs it if missing.
func checkAndInstall(cmd string, installCmd string) error {
	if _, err := exec.LookPath(cmd); err != nil {
		fmt.Printf("%s is not installed. Installing...\n", cmd)
		command := strings.Split(installCmd, " ")
		return exec.Command(command[0], command[1:]...).Run()
	}
	fmt.Printf("%s is already installed.\n", cmd)
	return nil
}

// installGo installs Go based on the detected OS/architecture.
func installGo() error {
	if _, err := exec.LookPath("go"); err == nil {
		fmt.Println("Go is already installed.")
		return nil
	}

	fmt.Println("Go is not installed. Installing...")
	var downloadURL string
	var installCmd string

	switch runtime.GOOS {
	case "linux":
		downloadURL = "https://dl.google.com/go/go1.23.5.linux-amd64.tar.gz"
		installCmd = "sudo tar -C /usr/local -xzf go1.23.5.linux-amd64.tar.gz && export PATH=$PATH:/usr/local/go/bin"
	case "darwin":
		downloadURL = "https://go.dev/dl/go1.23.5.darwin-amd64.pkg"
		installCmd = "sudo installer -pkg go1.23.5.darwin-amd64.pkg -target /"
	case "windows":
		downloadURL = "https://go.dev/dl/go1.23.5.windows-amd64.msi"
		installCmd = "msiexec /i go1.23.5.windows-amd64.msi /quiet /norestart"
	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	// Download Go using wget
	downloadCmd := fmt.Sprintf("wget %s", downloadURL)
	if err := exec.Command("sh", "-c", downloadCmd).Run(); err != nil {
		return fmt.Errorf("failed to download Go: %v", err)
	}

	// Install Go
	if err := exec.Command("sh", "-c", installCmd).Run(); err != nil {
		return fmt.Errorf("failed to install Go: %v", err)
	}

	fmt.Println("Go has been installed successfully.")
	return nil
}


// installTools installs the required tools.
func installTools() error {
	var pkgManager, updateCmd, installCmd string

	switch runtime.GOOS {
	case "linux":
		if _, err := exec.LookPath("apt-get"); err == nil {
			pkgManager = "apt-get"
			updateCmd = "sudo apt-get update -y && sudo apt-get upgrade -y"
			installCmd = "sudo apt-get install -y"
		} else if _, err := exec.LookPath("pacman"); err == nil {
			pkgManager = "pacman"
			updateCmd = "sudo pacman -Syu --noconfirm"
			installCmd = "sudo pacman -S --noconfirm"
		} else {
			return fmt.Errorf("unsupported Linux package manager")
		}
	case "darwin":
		pkgManager = "brew"
		updateCmd = "brew update && brew upgrade"
		installCmd = "brew install"
	 case "windows":
        tools := map[string]string{
            "curl": "winget install --id Microsoft.Curl",
            "git":  "winget install --id Git.Git",
            "npm":  "winget install --id OpenJS.NodeJS",
        }

        for cmd, installCmd := range tools {
            if _, err := exec.LookPath(cmd); err != nil {
                fmt.Printf("%s is not installed. Installing...\n", cmd)
                command := strings.Split(installCmd, " ")
                if err := exec.Command(command[0], command[1:]...).Run(); err != nil {
                    return fmt.Errorf("failed to install %s: %v", cmd, err)
                }
                fmt.Printf("%s installed successfully.\n", cmd)
            } else {
                fmt.Printf("%s is already installed.\n", cmd)
            }
        }
    default:
        return fmt.Errorf("tool installation is not supported on this platform: %s", runtime.GOOS)
	}

	fmt.Printf("Detected %s. Updating system...\n", pkgManager)
	if err := exec.Command("sh", "-c", updateCmd).Run(); err != nil {
		return fmt.Errorf("failed to update system: %v", err)
	}

	tools := map[string]string{
		"curl": "curl",
		"git":  "git",
		"zsh":  "zsh",
		"npm":  "npm",
		"tar":	"tar",
	}

	for cmd, pkg := range tools {
		if err := checkAndInstall(cmd, fmt.Sprintf("%s %s", installCmd, pkg)); err != nil {
			return fmt.Errorf("failed to install %s: %v", pkg, err)
		}
	}

	fmt.Println("All tools installed successfully.")
	return nil
}

func main() {
	// Install Go first
	if err := installGo(); err != nil {
		fmt.Fprintf(os.Stderr, "Error installing Go: %v\n", err)
		os.Exit(1)
	}

	// Install other tools
	if err := installTools(); err != nil {
		fmt.Fprintf(os.Stderr, "Error installing tools: %v\n", err)
		os.Exit(1)
	}
}
*/
