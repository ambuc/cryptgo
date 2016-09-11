package main

import "fmt"
import "errors"
import "bufio"
import "os"
import rand "math/rand"
import m "math"
import "time"
import "strings"
import "sort"
import "strconv"

type substitution struct {
	hint  string
	input string
	key   string
}

func (this substitution) encrypt() (string, error) {
	if this.key == "" {
		return "", errors.New("Malformed encryption: string <key> must be non-empty.")
	}

	return strings.Map(substitutionMapFunc(this.key, true), this.input), nil
}

func (this substitution) decrypt() (string, error) {
	switch this.hint {
	case "known":
		if this.key == "" {
			return "", errors.New("Malformed decryption: string <key> must be non-empty.")
		}
		return strings.Map(substitutionMapFunc(this.key, false), this.input), nil

	case "analyze":
		fmt.Println(generateGuessKey(pure(this.input)))
		key := generateEducatedGuess(this.input, 2000, false)
		return strings.Map(substitutionMapFunc(key, false), this.input), nil

	case "analyze-verbose":
		key := generateEducatedGuess(this.input, 2000, false)
		return strings.Map(substitutionMapFunc(key, false), this.input), nil
	}
	return "", errors.New("no hint given. specify `--hint known` or `--hint analyze` or `--hint analyize-verbose`")
}

///////////////////////
// LIBRARY FUNCTIONS //
///////////////////////

//implements hill-climbing algorithm
func generateEducatedGuess(input string, n int, verbose bool) string {
	quadgrams := getGrams("corpus/english_quadgrams.txt")

	pureInput := pure(input)

	evaluateKeyFitness := func(key string) float64 {
		return fitness(strings.Map(substitutionMapFunc(key, false), pureInput), quadgrams)
	}

	bestKey := generateGuessKey(pureInput)
	bestFitness := evaluateKeyFitness(bestKey)

	if verbose {
		fmt.Println(bestKey, bestFitness)
	}
	for i := 0; i < n; i++ {
		currentKey := mutateKey(bestKey)
		currentFitness := evaluateKeyFitness(currentKey)
		if currentFitness > bestFitness {
			bestKey = currentKey
			bestFitness = currentFitness
			if verbose {
				fmt.Println(bestKey, bestFitness)
			}
		}
	}
	return bestKey
}

func generateGuessKey(input string) string {
	n := make(map[float64]int)
	var vals []float64

	for k, v := range frequencyMap(input) {
		n[v] = k
		vals = append(vals, v)
	}

	sort.Sort(sort.Reverse(sort.Float64Slice(vals)))

	//e  t  a  o  i  n  s  h  r  d  l  u  c  m  f  w  y  p  v  b  g  k  j  q  x  z
	//5  20 1  15 9  14 19 8  18 4  12 21 3  13 6  23 25 16 22 2  7  11 10 17 24 26
	//a  b  c  d  e  f  g  h  i  j  k  l  m  n  o  p  q  r  s  t  u  v  w  x  y  z
	//1  2  3  4  5  6  7  8  9  10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26

	lookup := [26]int{5, 20, 1, 15, 9, 14, 19, 8, 18, 4, 12, 21, 3, 13, 6, 23, 25, 16, 22, 2, 7, 11, 10, 17, 24, 26}
	cipherAlphabet := [26]rune{}
	for i := 0; i < len(vals); i++ {
		cipherAlphabet[lookup[i]-1] = rune(n[vals[i]] + 65)
	}

	//fill out the rest of the cipheralphabet
	for runeIndex := 65; runeIndex < 65+26; runeIndex++ {
		flag := false
		for i := 0; i < 26; i++ {
			if int(cipherAlphabet[i]) == runeIndex {
				flag = true
			}
		}
		if !flag {
			for i := 0; i < 26; i++ {
				if cipherAlphabet[i] == rune(0) {
					cipherAlphabet[i] = rune(runeIndex)
					break
				}
			}
		}
	}
	s := ""
	for i := 0; i < 26; i++ {
		s = s + string(rune(cipherAlphabet[i]))
	}
	return s
}

func mutateKey(key string) string {
	newRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	q := sort.StringSlice(strings.Split(key, ""))
	q.Swap(newRand.Intn(26), newRand.Intn(26))
	return strings.Join(q, "")
}

func fitness(input string, quadgrams map[string]int) float64 {
	pureInput := strings.ToUpper(pure(input))

	var sum, n float64
	sum = 0
	n = 250000.0

	for i := 0; i < len(pureInput)-3; i++ {
		t := pureInput[i : i+4]
		if quadgrams[t] != 0 {
			q := m.Log(float64(quadgrams[t]) / n)
			sum = sum + q
		}
	}
	return sum
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

func getGrams(path string) map[string]int {
	lines, err := readLines(path)
	check(err)
	grams := make(map[string]int)
	for _, line := range lines {
		l := strings.Split(line, " ")
		grams[l[0]], _ = strconv.Atoi(l[1])
	}
	return grams
}
