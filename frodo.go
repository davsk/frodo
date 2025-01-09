// 'frodo.go'

// Package frodo automates updates.
package frodo // import "go.davsk.net/frodo"

import (
  runtime

  go.davsk.net/frodo/pkg/auto_update

)

func main() {

	auto_update.MustDoit("frodo","github.com/davsk/frodo","v0.0.1")

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
   } else if ( not pacman ) {
      ReportError("Not Arch Linux"
                  } else {
                  ArchMustDoit()
   }
  }
