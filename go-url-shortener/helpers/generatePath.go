package helpers

import (
	"math/rand"
	"time"
)

var source = rand.NewSource(time.Now().UnixNano())

const _charsetLowercase = "abcdefghijklmnoprstuvwyz"
const _charsetUppercase = "ABCDEFGHIJKLMNOPRSTUVWYZ"
const _charsetNumbers = "1234567890"

// Options ...
type Options struct {
	Length    int
	Uppercase bool
	Lowercase bool
	Numbers   bool
}

// GeneratePath is a function that makes a new random path.
func GeneratePath() string {
	x := make([]byte, 10)

	charset := "."

	opt := Options{
		10,
		true,
		true,
		true,
	}

	if opt.Uppercase {
		charset += _charsetUppercase
	}
	if opt.Lowercase {
		charset += _charsetLowercase
	}
	if opt.Numbers {
		charset += _charsetNumbers
	}

	for i := range x {
		x[i] = charset[source.Int63()%int64(len(charset))]
	}

	return string(x)
}
