package main

import "fmt"
import m "math"
import "strings"
import "regexp"

func check(e error) {
  if e != nil {
    panic(e)
  }
}



func shiftChar(r rune, shift int) rune {
  //fmt.Println(string(r), r, int(r)+shift, string(int(r)+shift))
  if( 65<=r && r<=90 ) {
    return rune((((int(r) - 65 ) + shift) % 26) + 65)
  }
  if( 97<=r && r<=122 ) {
    return rune((((int(r) - 97 ) + shift) % 26) + 97)
  }
  return r
}

func shiftWord(inputText string, n int) string {
  return strings.Map( func (r rune) rune { return shiftChar(r, n) }, inputText)
}

func strip(s string) string {
  return regexp.MustCompile("[^a-zA-Z]").ReplaceAllString(s, "")
}

func shorten(s string) string {
  result := regexp.MustCompile("\n").ReplaceAllString(s,"")
  if(len(s) > 50){
    return result[:45] + "..."
  }
  return result + "..."
}


func pure(input string) string {
  return strings.ToLower(strip(input))
}

func freq(input string) map[int]float64 {
  testMap := make(map[int]float64)
  for _, c := range input {
    testMap[int(c)-97] = testMap[int(c)-97] + 1
  }
  for k, _ := range testMap {
    testMap[k] = testMap[k] / float64(len(input))
  }
  return testMap
}

func dist(a []float64, b []float64) float64 {
  i := 0
  sum := 0.0
  for (i < len(a)) {
    sum = sum + (a[i] - b[i])*(a[i] - b[i])
    i = i + 1
  }
  return m.Sqrt(sum)
}

func mapToArray(input map[int]float64) []float64 {
  vals := make([]float64, 0, 26)
  i := 0
  for (i < 26) {
    vals = append(vals, input[i])
    i = i + 1
  }
  return vals
}

func english() map[int]float64 {
  return map[int]float64{ 0:0.08167,  1:0.01492,  2:0.02782,  3:0.04253,  4:0.12702, 5:0.02228,  6:0.02015,  7:0.06094,  8:0.06966,  9:0.00153, 10:0.00772, 11:0.04025, 12:0.02406, 13:0.06749, 14:0.07507, 15:0.01929, 16:0.00095, 17:0.05987, 18:0.06327, 19:0.09056, 20:0.02758, 21:0.00978, 22:0.02360, 23:0.00150, 24:0.01974, 25:0.00074 } 
}

func frequencyAnalysis(input string, verbose bool) string {

  cleanInput := pure(input)

  engArr := mapToArray(english())

  i := 0
  poss := map[int]float64{}
  for (i < 26) {
    shifted := shiftWord(cleanInput, i)
    testArr := mapToArray(freq(shifted))
    poss[i] = dist(engArr, testArr)
    i = i + 1
  }

  globalMin := 1000.0
  bestKey := 0
  for k, v := range poss {
    if (v < globalMin){
      globalMin = v; bestKey = k
    }
  }

  if (verbose) {
    for {
      if (len(poss) == 0) { break }
      localMax := 0.0
      localMaxKey := 0 
      for k,v := range poss {
        if v >= localMax {
          localMax = v
          localMaxKey = k
        }
      }
      fmt.Printf(" %3f :: %v\n", localMax, shorten(shiftWord(input, localMaxKey)))
      delete(poss, localMaxKey)
    }
    fmt.Printf("(lower is better)\n")
  }

  return shiftWord(input, bestKey)
}
