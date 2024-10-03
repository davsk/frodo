// 'frodo.go'

// Package frodo automates updates.
package frodo // import "go.davsk.net/frodo"

import (
  runtime
)

isDarwin := false
isLinux := false
isWindows := false
isUnknown := false

func main() {
  switch ( runtime.GOOS )
  case "darwin"
    isDarwin = true
  case "linux"
   isLinux = true
  case "windows"
    isWindows = true
  default
    isUnknown = true
}
