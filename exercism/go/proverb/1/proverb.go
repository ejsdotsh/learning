// Package proverb generates a proverb in the form of the saying,
// "For want of a horseshoe nail, a kingdom was lost"
package proverb

import "fmt"

// Proverb takes a list of inputs and generates the relevant proverb.
func Proverb(rhyme []string) []string {
	singleInput := "And all for the want of a %s."
	multiInput := "For want of a %s the %s was lost."
	res := []string{}

	if len(rhyme) > 0 {
		for i := range len(rhyme) - 1 {
			res = append(res, fmt.Sprintf(multiInput, rhyme[i], rhyme[i+1]))
		}
		res = append(res, fmt.Sprintf(singleInput, rhyme[0]))
	}

	return res
}
