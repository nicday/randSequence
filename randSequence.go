// Package randSequence contains a utility function for generating random sequences.
package randSequence

import "math/rand"

var (
	// DefaultCharacters is the default characters for a new RandGenerator.
	DefaultCharacters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
)

// RandGenerator represents a random sequence generator.
type RandGenerator struct {
	// Characters is a slice of runes that is used to generate the random sequence.
	Characters []rune
}

// New creates a new RandomGenerator initialized with lowercase, uppercase and numeric characters.
func New() *RandGenerator {
	return &RandGenerator{
		Characters: DefaultCharacters,
	}
}

// Generate returns a new random sequence of length `n`.
func (r RandGenerator) Generate(n int) string {
	s := make([]rune, n)
	for i := range s {
		s[i] = r.Characters[rand.Intn(len(r.Characters))]
	}
	return string(s)
}
