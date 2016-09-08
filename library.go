package main

import "fmt"
import m "math"
import "strings"
import "regexp"
import "os"
import "text/tabwriter"

func check(e error) {
  if e != nil {
    panic(e)
  }
}

//euclidean algorithm for coprimity
func GCDIterative(u, v int) bool {
  var t int
  for u > 0 {
    if u < v {
      t = u
      u = v
      v = t
    }
    u = u - v
  }
  return (v == 1)
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
  result := regexp.MustCompile("\n").ReplaceAllString(s,"")
  if(len(s) > 50){
    return result[:45] + "..."
  }
  return result + "..."
}

//returns a lowercase alphanumeric string without spaces
func pure(input string) string {
  return strings.ToLower(strip(input))
}

///////////////////////////
// FANCY STRING PRINTING //
///////////////////////////

// prints a given string at a list of [int] shifts, ranked by (float64)s in (in|de)creasing order
func verboselyPrintByScore(poss map[int]float64, input string, decreasing bool){
  p := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)
  if (decreasing) {
    for {
      if (len(poss) == 0) { break }
      localMax := 0.0; localMaxKey := 0 
      for k,v := range poss {
        if v >= localMax {
          localMax = v; localMaxKey = k
        }
      }
      fmt.Fprintf(p, "+%v \t%3f \t %v\n", byte(localMaxKey), localMax, shorten(shiftWord(input, localMaxKey)))
      delete(poss, localMaxKey)
    }

  } else {

    for {
      if (len(poss) == 0) { break }
      localMin := 100.0; localMinKey := 0
      for k,v := range poss {
        if v <= localMin {
          localMin = v; localMinKey = k
        }
      }
      fmt.Fprintf(p, "+%v \t%3f \t %v\n", byte(localMinKey), localMin, shorten(shiftWord(input, localMinKey)))
      delete(poss, localMinKey)
    }
  }

  p.Flush()

  if (decreasing) {
    fmt.Printf("Scores ranked in decreasing order. (Lower is better)\n")
  } else {
    fmt.Printf("Scores ranked in increasing order. (Higher is better)\n")
  }
}

/////////////////////////////
// CRYPTOGRAPHIC FUNCTIONS //
/////////////////////////////

// shifts a single character by n
func shiftChar(r rune, shift int) rune {
  if (shift < 0) { shift = (shift + 26) % 26 }
  if( 65<=r && r<=90 ) {
    return rune((((int(r) - 65 ) + shift) % 26) + 65)
  }
  if( 97<=r && r<=122 ) {
    return rune((((int(r) - 97 ) + shift) % 26) + 97)
  }
  return r
}

// shifts an entire word by n characters
func shiftWord(inputText string, n int) string {
  return strings.Map( func (r rune) rune { return shiftChar(r, n) }, inputText)
}

// return a * (r + b ) + c
func affineShift(a int, b int, c int) func(r rune) rune {
  return func(r rune) rune {
    if( 65<=r && r<=90 ) {
      return rune(((a*(int(r)-65+b)+c)%26+26)%26+65)
    } else if( 97<=r && r<=122 ) {
      return rune(((a*(int(r)-97+b)+c)%26+26)%26+97)
    }
    return r
  }
}

// for use in the atbash cipher
func flip(r rune) rune {
  switch {
  case 65 <= r && r <= 90:
    return (26-(r-'A')) + 'A' - 1
  case 97 <= r && r <= 122:
    return (26-(r-'a')) + 'a' - 1
  }
  return r
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

//given two lists of float64s, returns a distance, where lower is better
func euclideanDistance(a []float64, b []float64) float64 {
  i := 0
  sum := 0.0
  for (i < len(a)) {
    sum = sum + (a[i] - b[i])*(a[i] - b[i])
    i = i + 1
  }
  return m.Sqrt(sum)
}

//given a map of float64s, returns a list of values
func mapToArray(input map[int]float64) []float64 {
  vals := make([]float64, 0, 26)
  i := 0
  for (i < 26) {
    vals = append(vals, input[i])
    i = i + 1
  }
  return vals
}

// returns a map of english letter frequencies
func english() map[int]float64 {
  return map[int]float64{ 0:0.08167,  1:0.01492,  2:0.02782,  3:0.04253,  4:0.12702, 5:0.02228,  6:0.02015,  7:0.06094,  8:0.06966,  9:0.00153, 10:0.00772, 11:0.04025, 12:0.02406, 13:0.06749, 14:0.07507, 15:0.01929, 16:0.00095, 17:0.05987, 18:0.06327, 19:0.09056, 20:0.02758, 21:0.00978, 22:0.02360, 23:0.00150, 24:0.01974, 25:0.00074 } 
}

//performs frequency analysis on an input string
func frequencyAnalysis(input string, verbose bool) string {
  poss := map[int]float64{}
  i := 0
  for (i < 26) {
    poss[i] = euclideanDistance( mapToArray( english()), mapToArray( frequencyMap( shiftWord(pure(input), i))))
    i = i + 1
  }

  globalMin := 1000.0; globalMinKey := 0
  for k, v := range poss {
    if (v < globalMin){
      globalMin = v; globalMinKey = k
    }
  }

  if (verbose) { verboselyPrintByScore(poss, input, true) }

  return shiftWord(input, globalMinKey)
}

