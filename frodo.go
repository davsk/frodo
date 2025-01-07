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
isAmd64 := false

func main() {
  switch ( runtime.GOOS ) {
  case "darwin":
    DarwinMustDoit( runtime.GOARCH )
  case "linux":
    LinuxMustDoit( runtime.GOARCH )
  case "windows":
    WindowsMustDoit( runtime.GOARCH )
  default:
    ReportUnknown( runtime.GOOS, runtime.GOARCH )
}
  )

func DarwinMustDoit(arch) { 
  //
}  

 func LinuxMustDoit(arch) {
   if ( arch != "Amd86" ) {
     ReportUnknown( runtime.GOOS, runtime,GOARCH )
   } else if ( // test for pacman ) {
   }
  
  switch ( runtime.GOARCH ) {
    case "amd64" :
    isAmd64 = true
    default:
    isAmd64 = false
    }

  if ( isDarwin ) {
    Darwin.mustDoit()
    } else if ( isLinux 
  }
