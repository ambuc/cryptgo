package main

import "strings"

type atbash struct {
	input string
}

func (this atbash) encrypt() (string, error) {
	return strings.Map(atbashFlip, this.input), nil
}

func (this atbash) decrypt() (string, error) {
	return strings.Map(atbashFlip, this.input), nil
}

///////////////////////
// LIBRARY FUNCTIONS //
///////////////////////

// for use in the atbash cipher
func atbashFlip(r rune) rune {
	switch {
	case 65 <= r && r <= 90:
		return (26 - (r - 'A')) + 'A' - 1
	case 97 <= r && r <= 122:
		return (26 - (r - 'a')) + 'a' - 1
	}
	return r
}
