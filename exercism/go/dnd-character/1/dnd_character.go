package dndcharacter

import (
	"math"
	"math/rand/v2"
	"slices"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor((float64(score) - 10) / 2))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	scores := make([]int, 0)
	minScore := 1
	maxScore := 7 // really max + 1
	for range 4 {
		s := rand.IntN(maxScore-minScore) + minScore
		scores = append(scores, s)
	}
	slices.Sort(scores)
	score := scores[1] + scores[2] + scores[3]

	return score
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	ch := Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: Ability(),
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
		Hitpoints:    10,
	}
	ch.Hitpoints = ch.Hitpoints + Modifier(ch.Constitution)

	return ch
}
