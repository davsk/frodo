// 'cmd/frodo/frodo.go'

package main

import (
	//"fmt"

	//"go.davsk.net/frodo/pkg/auto_update"
	//"go.davsk.net/frodo/pkg/config"
	//"go.davsk.net/frodo/pkg/install"
	"go.davsk.net/frodo/pkg/privilege"
)

func main() {
	// Load profile from environment and gh
	//profile := config.Load()

	//auto_update.MustDoit("frodo", "github.com/davsk/frodo", "v0.1.6-alpha")

	privilege.Doit("yay")
	//fmt.Println(a)
	//	update.System()
	//} else {
	//	install.Hashicorp()
	//}
}
