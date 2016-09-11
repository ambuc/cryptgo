package main

import "bufio"
import "os"
import m "math"
import "regexp"
import "strconv"
import "strings"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// returns a map of english letter frequencies
func english() map[int]float64 {
	return map[int]float64{0: 0.08167, 1: 0.01492, 2: 0.02782, 3: 0.04253, 4: 0.12702, 5: 0.02228, 6: 0.02015, 7: 0.06094, 8: 0.06966, 9: 0.00153, 10: 0.00772, 11: 0.04025, 12: 0.02406, 13: 0.06749, 14: 0.07507, 15: 0.01929, 16: 0.00095, 17: 0.05987, 18: 0.06327, 19: 0.09056, 20: 0.02758, 21: 0.00978, 22: 0.02360, 23: 0.00150, 24: 0.01974, 25: 0.00074}
}

//euclidean algorithm for coprimity
func GCDIterative(u, v int) bool {
	for u > 0 {
		if u < v {
			t := u
			u = v
			v = t
		}
		u = u - v
	}
	return (v == 1)
}

//given two lists of float64s, returns a distance
func euclideanDistance(a []float64, b []float64) float64 {
	i := 0
	sum := 0.0
	for i < len(a) {
		sum = sum + (a[i]-b[i])*(a[i]-b[i])
		i = i + 1
	}
	return m.Sqrt(sum)
}

///////////////////
// MAP FUNCTIONS //
///////////////////

func invertRuneMap(m map[rune]rune) map[rune]rune {
	n := make(map[rune]rune)
	for k, v := range m {
		n[v] = k
	}
	return n
}

//given a map of float64s, returns a list of values
func getFloat64MapVals(input map[int]float64) []float64 {
	vals := make([]float64, 0, 26)
	i := 0
	for i < 26 {
		vals = append(vals, input[i])
		i = i + 1
	}
	return vals
}

//given an input, returns a map of characters (by int) and their frequencies (by float64)
func frequencyMap(input string) map[int]float64 {
	testMap := make(map[int]float64)
	for _, c := range input {
		testMap[int(c)-97] = testMap[int(c)-97] + 1
	}
	for k, _ := range testMap {
		testMap[k] = testMap[k] / float64(len(input))
	}
	return testMap
}

//////////////////////
// STRING FUNCTIONS //
//////////////////////

//returns a mixed-case alphanumeric string without spaces
func strip(s string) string {
	return regexp.MustCompile("[^a-zA-Z]").ReplaceAllString(s, "")
}

//trims a string to the first 50 characters and removes leading/trailing newlines
func shorten(s string) string {
	result := regexp.MustCompile("\n").ReplaceAllString(s, "")
	if len(s) > 50 {
		return result[:45] + "..."
	}
	return result + "..."
}

//returns a lowercase alphanumeric string without spaces
func pure(input string) string {
	return strings.ToLower(strip(input))
}

/////////////////////////////
// CRYPTOGRAPHIC FUNCTIONS //
/////////////////////////////

// shifts a single character by n
func shiftChar(r rune, shift int) rune {
	if shift < 0 {
		shift = (shift + 26) % 26
	}
	if 65 <= r && r <= 90 {
		return rune((((int(r) - 65) + shift) % 26) + 65)
	}
	if 97 <= r && r <= 122 {
		return rune((((int(r) - 97) + shift) % 26) + 97)
	}
	return r
}

// shifts an entire word by n characters
func shiftWord(inputText string, n int) string {
	return strings.Map(func(r rune) rune { return shiftChar(r, n) }, inputText)
}

///////////////////
// I/O FUNCTIONS //
///////////////////

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func readCorpus(path string) map[string]int {
	lines, err := readLines(path)
	check(err)
	grams := make(map[string]int)
	for _, line := range lines {
		l := strings.Split(line, " ")
		grams[l[0]], _ = strconv.Atoi(l[1])
	}
	return grams
}
