// go.davsk.net/frodo/pkg/auto_update

// package auto_update is a davsk wrapper for go-rocket-update
package auto_update // import "go.davsk.net/frodo/pkg/auto_update"

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"sync"

	"github.com/mouuff/go-rocket-update/pkg/provider"
	"github.com/mouuff/go-rocket-update/pkg/updater"
)

// func MustDoit checks for a new version and updates when one is found.
// It recieves prgName, repoURL, and semVer.
func MustDoit(prgName string, repoURL string, semVer string) {
	// Determine the appropriate file name extension based on the operating system
	var fileNameExtension string
	switch runtime.GOOS {
	case "windows":
		fileNameExtension = "zip"
	default:
		fileNameExtension = "tar.gz"
	}

	// define updater
	u := &updater.Updater{
		Provider: &provider.Github{
			RepositoryURL: repoURL,
			ArchiveName:   fmt.Sprintf("frodo_%s.%s", runtime.GOOS, fileNameExtension),
		},
		ExecutableName: fmt.Sprintf("%s_%s_%s", prgName, runtime.GOOS, runtime.GOARCH),
		Version:        semVer, // You can change this value to trigger an update
	}

	versionFlag := false
	flag.BoolVar(&versionFlag, "version", false, "prints the version and exit")
	flag.Parse()

	if versionFlag {
		// we use this flag to verify the installation for this example:
		// https://github.com/mouuff/go-rocket-update/blob/master/examples/github-rollback/main.go
		fmt.Println(u.Version)
		return
	}

	log.Println("Current version: " + u.Version)
	log.Println("Looking for updates...")
	var wg sync.WaitGroup
	wg.Add(1)
	// For the example we run the update in the background
	// but you could directly call u.Update()
	var updateStatus updater.UpdateStatus
	var updateErr error
	go func() {
		updateStatus, updateErr = u.Update()
		wg.Done()
	}()

	// Here you can execute your program

	wg.Wait() // Waiting for the update process to finish before exiting
	if updateErr != nil {
		log.Println(updateErr)
	}
	if updateStatus == updater.Updated {
		log.Println("Updated!")
	}
}
