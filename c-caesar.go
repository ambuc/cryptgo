package main

import "errors"
import "fmt"
import "os"
import "text/tabwriter"

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
	case "known":
		return shiftWord(c.input, -c.n), nil
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
		poss[i] = euclideanDistance(mapToArray(english()), mapToArray(frequencyMap(shiftWord(pure(input), i))))
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
		caesarVerboselyPrintByScore(poss, input, true)
	}

	return shiftWord(input, globalMinKey)
}

// prints a given string at a list of [int] shifts, ranked by (float64)s in (in|de)creasing order
func caesarVerboselyPrintByScore(poss map[int]float64, input string, decreasing bool) {
	p := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)
	if decreasing {
		for {
			if len(poss) == 0 {
				break
			}
			localMax := 0.0
			localMaxKey := 0
			for k, v := range poss {
				if v >= localMax {
					localMax = v
					localMaxKey = k
				}
			}
			fmt.Fprintf(p, "+%v \t%3f \t %v\n", byte(localMaxKey), localMax, shorten(shiftWord(input, localMaxKey)))
			delete(poss, localMaxKey)
		}

	} else {

		for {
			if len(poss) == 0 {
				break
			}
			localMin := 100.0
			localMinKey := 0
			for k, v := range poss {
				if v <= localMin {
					localMin = v
					localMinKey = k
				}
			}
			fmt.Fprintf(p, "+%v \t%3f \t %v\n", byte(localMinKey), localMin, shorten(shiftWord(input, localMinKey)))
			delete(poss, localMinKey)
		}
	}

	p.Flush()

	if decreasing {
		fmt.Printf("Scores ranked in decreasing order. (Lower is better)\n")
	} else {
		fmt.Printf("Scores ranked in increasing order. (Higher is better)\n")
	}
}
