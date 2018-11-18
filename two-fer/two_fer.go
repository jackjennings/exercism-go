// Package twofer implements functions about sharing.
package twofer

import "fmt"

// ShareWith returns a phrase describing sharing two of something with someone.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
