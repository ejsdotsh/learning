// Package twofer returns a custom greeting if given a name.
package twofer

import "fmt"

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	who := "you"
	if name != "" {
		who = name
	}
	return fmt.Sprintf("One for %s, one for me.", who)
}
