// Package config contains the init, edit, save, and load of the frodo config.
package config

import "fmt"

type Profile struct {
	UserName, UserEmail string
}

// New
func Load() *Profile {
	fmt.Printf("test.")
	return new(Profile)
}
