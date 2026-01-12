// Package twofer provides cookie sharing messages.
package twofer

import "fmt"

// ShareWith returns a sharing message for the given name.
// If name is empty, "you" is used instead.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}