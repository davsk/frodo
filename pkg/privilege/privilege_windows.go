// +build windows
// file: 'frodo/pkg/privilege/privilege_windows.go'


// by David Lynn Skinner
// on January 16, 2025
// for Davsk Ltd Co

// copyright (c) 2025 Davsk Ltd Co

// package privilege checks for root, sudo, or admin privilege
// and restarts with the requested privilege

package privilege

import (
   	"syscall"
   	"time"
   	"os"
   	"strings"
   	"fmt"
   	"golang.org/x/sys/windows"
)

func becomeAdmin() {
   	verb := "runas"
   	exe, _ := os.Executable()
   	cwd, _ := os.Getwd()
   	args := strings.Join(os.Args[1:], " ")

   	verbPtr, _ := syscall.UTF16PtrFromString(verb)
   	exePtr, _ := syscall.UTF16PtrFromString(exe)
   	cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
   	argPtr, _ := syscall.UTF16PtrFromString(args)

   	var showCmd int32: 1 //SW_NORMAL

   	err := windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
   	if err != nil {
  		fmt.Println(err)
   	}
}

