package main

import "errors"
import m "math"

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
		return caesarFrequencyAnalysis(c.input, false), nil
	case "analyze-verbose":
		return caesarFrequencyAnalysis(c.input, true), nil
	}
	return "", errors.New("no hint given. specify `--hint brute-force` or `--hint analyze` or `--hint analyize-verbose`")
}

///////////////////////
// LIBRARY FUNCTIONS //
///////////////////////

//performs frequency analysis on an input string
func caesarFrequencyAnalysis(input string, verbose bool) string {
	poss := map[int]float64{}
	i := 0
	for i < 26 {
		poss[i] = euclideanDistance(mapToArray(english()), mapToArray(casearFrequencyMap(shiftWord(pure(input), i))))
		i = i + 1
	}

	globalMin := 1000.0
	globalMinKey := 0
	for k, v := range poss {
		if v < globalMin {
			globalMin = v
			globalMinKey = k
		}
	}

	if verbose {
		verboselyPrintByScore(poss, input, true)
	}

	return shiftWord(input, globalMinKey)
}

//given an input, returns a map of characters (by int) and their frequencies (by float64)
func caesarFrequencyMap(input string) map[int]float64 {
	testMap := make(map[int]float64)
	for _, c := range input {
		testMap[int(c)-97] = testMap[int(c)-97] + 1
	}
	for k, _ := range testMap {
		testMap[k] = testMap[k] / float64(len(input))
	}
	return testMap
}
