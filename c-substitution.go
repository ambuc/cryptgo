package main

import "errors"
import "fmt"
import "os"
import m "math"
import rand "math/rand"
import "sort"
import "strings"
import "text/tabwriter"
import "time"

type substitution struct {
	hint  string
	input string
	key   string
	n     int
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
		key := substitutionHillClimb(this.input, this.n, false)
		return strings.Map(substitutionMapFunc(key, false), this.input), nil

	case "analyze-verbose":
		key := substitutionHillClimb(this.input, this.n, true)
		return strings.Map(substitutionMapFunc(key, false), this.input), nil
	}
	return "", errors.New("no hint given. specify `--hint known` or `--hint analyze` or `--hint analyize-verbose`")
}

///////////////////////
// LIBRARY FUNCTIONS //
///////////////////////

//implements hill-climbing algorithm
func substitutionHillClimb(dirtyInput string, n int, verbose bool) string {
	var iterations int
	if n != 0 {
		iterations = n
	} else {
		iterations = 2000
	}
	quadgrams := readCorpus("corpus/english_quadgrams.txt")

	input := pure(dirtyInput)

	//given some text, return an arbitrary value for how english-like it is.
	getFitness := func(dirtyText string, quadgrams map[string]int) float64 {
		text := strings.ToUpper(dirtyText)
		sum := 0.0
		for i := 0; i < len(text)-3; i++ {
			t := text[i : i+4]
			if quadgrams[t] != 0 {
				sum = sum + m.Log(float64(quadgrams[t])/250000.0)
			}
		}
		return sum
	}

	//given some key, return a fitness value
	//for how a substitution decryption performs under it
	evaluate := func(key string) float64 {
		dirtyText := strings.Map(substitutionMapFunc(key, false), input)
		return getFitness(dirtyText, quadgrams)
	}

	//given some key, return a copy with two random letters swapped
	mutate := func(key string) string {
		newRand := rand.New(rand.NewSource(time.Now().UnixNano()))
		q := sort.StringSlice(strings.Split(key, ""))
		q.Swap(newRand.Intn(26), newRand.Intn(26))
		return strings.Join(q, "")
	}

	bestKey := substitutionSeedKey(input)
	bestFitness := evaluate(bestKey)

	t := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)
	if verbose {
		fmt.Fprintln(t, "\nKey \t Fitness")
		fmt.Fprintf(t, "%v \t %2f \n", bestKey, bestFitness)
	}

	for i := 0; i < iterations; i++ {
		currentKey := mutate(bestKey)
		currentFitness := evaluate(currentKey)
		if currentFitness > bestFitness {
			bestKey = currentKey
			bestFitness = currentFitness
			if verbose {
				fmt.Fprintf(t, "%v \t %2f \n", bestKey, bestFitness)
			}

		}
	}
	if verbose {
		t.Flush()
	}
	return bestKey
}

func substitutionSeedKey(input string) string {

	mapFrequenciesToLetterIndices := make(map[float64]int)

	var vals []float64

	//generate a mapping of frequencies to letter indices
	for k, v := range frequencyMap(input) {
		mapFrequenciesToLetterIndices[v] = k
		vals = append(vals, v)
	}

	//sort the frequencies in decreasing order
	sort.Sort(sort.Reverse(sort.Float64Slice(vals)))

	//hard-coded map corresponding to the letter indices of etaoin shrdlu
	lookup := [26]int{5, 20, 1, 15, 9, 14, 19, 8, 18, 4, 12, 21, 3, 13, 6, 23, 25, 16, 22, 2, 7, 11, 10, 17, 24, 26}

	//generate a cipherAlphabet from the most popular letters we do have
	cipherAlphabet := [26]rune{}
	for i := 0; i < len(vals); i++ {
		cipherAlphabet[lookup[i]-1] = rune(mapFrequenciesToLetterIndices[vals[i]] + 65)
	}

	//fill in the gaps in the cipherAlphabet for the letters we don't have
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

	//stitch it all together into a cipherAlphabet string
	s := make([]string, 26)
	for idx, val := range cipherAlphabet {
		s[idx] = string(rune(val))
	}
	return strings.Join(s, "")
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
