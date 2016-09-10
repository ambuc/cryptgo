package main

import "errors"

type caesar struct {
	hint  string
	input string
 	n     int
}

func (c caesar) encrypt() (string, error) {
	return shiftWord(c.input, c.n), nil
}

func (c caesar) decrypt() (string, error) {
	switch c.hint {
	case "brute-force":
		result := "\n"
		i := 0
		for i < 26 {
			result = result + shiftWord(c.input, i) + "\n"
			i = i + 1
		}
		return result, nil
	case "analyze":
		return frequencyAnalysis(c.input, false), nil
	case "analyze-verbose":
		return frequencyAnalysis(c.input, true), nil
	}
	return "", errors.New("no hint given. specify `--hint brute-force` or `--hint analyze` or `--hint analyize-verbose`")
}
