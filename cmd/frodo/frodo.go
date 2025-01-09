// 'cmd/frodo/frodo.go'

package main

import (
	"go.davsk.net/frodo/pkg/auto_update"
	"go.davsk.net/frodo/pkg/config"
)

func main() {
	auto_update.MustDoit("frodo", "github.com/davsk/frodo", "v0.0.1")
	config.Load()
}
