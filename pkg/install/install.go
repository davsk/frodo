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

// func hashicorp installs the hashicorp tools
func hashicorp() {
	cmd := exec.Command("brew", "update")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	must(err)
	cmd.Wait()
}
