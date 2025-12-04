// Package proverb generates a proverb in the form of the saying,
// "For want of a horseshoe nail, a kingdom was lost"
package proverb

import "fmt"

// Proverb takes a list of inputs and generates the relevant proverb.
func Proverb(rhyme []string) []string {
	// immediately return with no input
	if len(rhyme) == 0 {
		return []string{}
	}

	prvb := []string{}

	for i := range len(rhyme) - 1 {
		prvb = append(prvb, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
	}
	prvb = append(prvb, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))

	return prvb
}
