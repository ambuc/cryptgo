package main

//import "fmt"
import "errors"
import "strings"

//import "sort"

type substitution struct {
	hint  string
	input string
	key   string
}

func (this substitution) encrypt() (string, error) {
	if this.key == "" {
		return "", errors.New("Malformed encryption: string <key> must be non-empty.")
	}

	return strings.Map(substitutionMapFunc(this.key, false), this.input), nil
}

func (this substitution) decrypt() (string, error) {
	switch this.hint {
	case "known":
		if this.key == "" {
			return "", errors.New("Malformed decryption: string <key> must be non-empty.")
		}

		return strings.Map(substitutionMapFunc(this.key, true), this.input), nil

	case "analyze":
		return this.input, nil

	case "analyze-verbose":
		return this.input, nil
	}
	return "", errors.New("no hint given. specify `--hint known` or `--hint analyze` or `--hint analyize-verbose`")
}

func invertRuneMap(m map[rune]rune) map[rune]rune {
	n := make(map[rune]rune)
	for k, v := range m {
		n[v] = k
	}
	return n
}

func substitutionMap(key string) map[rune]rune {

	cipherMap := make(map[rune]rune)

	expandedAlphabet := strings.ToLower(key) + "abcdefghijklmnopqrstuvwxyz"
	uppercaseExpandedAlphabet := strings.ToUpper(expandedAlphabet)
	c, d := 0, 0
	for i := 0; i < len(key)+26; i++ {
		if cipherMap[rune(expandedAlphabet[i])] == 0 {
			cipherMap[rune(expandedAlphabet[i])] = rune(c + 97)
			c = c + 1
		}
		if cipherMap[rune(uppercaseExpandedAlphabet[i])] == 0 {
			cipherMap[rune(uppercaseExpandedAlphabet[i])] = rune(d + 65)
			d = d + 1
		}
	}
	return cipherMap
}

func substitutionMapFunc(key string, invert bool) func(r rune) rune {

	cipherMap := substitutionMap(key)
	if invert {
		cipherMap = invertRuneMap(cipherMap)
	}

	return func(r rune) rune {
		if 65 <= r && r <= 90 {
			return cipherMap[r]
		} else if 97 <= r && r <= 122 {
			return cipherMap[r]
		}
		return r
	}
}
