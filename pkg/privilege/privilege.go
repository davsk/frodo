// file: 'frodo/pkg/privilege/privilege.go'


// by David Lynn Skinner
// on January 10, 2025
// for Davsk Ltd Co

// copyright (c) 2025 Davsk Ltd Co

// package privilege checks for root, sudo, or admin privilege
// and restarts with the requested privilege
package privilege

import (
  "os"
  "time"
)

func checkAdmin() bool {
   	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")

    return err == nil
}

func Verify() {
  if !checkAdmin() {
  	becomeAdmin()

   	time.Sleep(2 * time.Second)

 	  os.Exit(0)
  }
}
