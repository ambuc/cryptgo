package main

import "fmt"
import "errors"
import "strings"
import "sort"

type affine struct {
	hint  string
	input string
	a     int
	b     int
}

func (this affine) encrypt() (string, error) {
	if this.a <= 0 {
		return "", errors.New("Insecure encryption: value <a> must be greater than zero")
	}
	if this.b <= 0 {
		return "", errors.New("Insecure encryption: value <b> must be greater than zero")
	}
	if !GCDIterative(26, this.a) {
		return "", errors.New("Insecure encryption: value <a> not coprime to 26")
	}
	return strings.Map(affineShift(this.a, 0, this.b), this.input), nil
}

func (this affine) decrypt() (string, error) {
	switch this.hint {
	case "known":
		if this.a <= 0 {
			return "", errors.New("Insecure encryption: value <a> must be greater than zero")
		}
		if this.b <= 0 {
			return "", errors.New("Insecure encryption: value <b> must be greater than zero")
		}
		if !GCDIterative(26, this.a) {
			return "", errors.New("Unreal decryption: <a> not coprime to 26")
		}
		return strings.Map(affineShift(affineModInverse(this.a), -this.b, 0), this.input), nil

	case "analyze":
		return affineAnalyze(this.input, false)

	case "analyze-verbose":
		return affineAnalyze(this.input, true)
	}
	return "", errors.New("no hint given. specify `--hint known` or `--hint analyze` or `--hint analyize-verbose`")
}

func affineAnalyze(input string, verbose bool) (string, error) {
	m := frequencyMap(strings.ToLower(pure(input)))
	mInv := make(map[float64]int)
	var frequencies []float64
	for letter, frequency := range m {
		frequencies = append(frequencies, frequency)
		mInv[frequency] = letter
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(frequencies)))
	p := 4                    //e
	q := 19                   //t
	r := mInv[frequencies[0]] //first most frequently occuring letter
	s := mInv[frequencies[1]] //second most frequently occuring letter
	D := ((p - q) + 26) % 26
	Dinv := affineModInverse(D)
	a := ((Dinv * (r - s) % 26) + 26) % 26
	b := ((Dinv * (p*s - q*r) % 26) + 26) % 26
	if verbose {
		fmt.Println("")
		fmt.Println("Two most frequent letters in English:")
		fmt.Println("\t e #4  -> p = 4")
		fmt.Println("\t t #19 -> q = 19")
		fmt.Println("")
		fmt.Println("Two most frequent letters in Ciphertext:")
		fmt.Printf("\t %v #%v -> r = %v\n", string(rune(r+97)), r, r)
		fmt.Printf("\t %v #%v -> s = %v\n", string(rune(s+97)), s, s)
		fmt.Println("")
		fmt.Println("Solving system of linear equations:")
		fmt.Println("\t a x p + b = r")
		fmt.Println("\t a x q + b = s")
		fmt.Println("")
		fmt.Println("Results:")
		fmt.Printf("\t a = %v; b = %v", a, b)
		fmt.Println("")
	}
	return strings.Map(affineShift(affineModInverse(a), -b, 0), input), nil
}

///////////////////////
// LIBRARY FUNCTIONS //
///////////////////////

// return a * (r + b ) + c
func affineShift(a int, b int, c int) func(r rune) rune {
	return func(r rune) rune {
		if 65 <= r && r <= 90 {
			return rune(((a*(int(r)-65+b)+c)%26+26)%26 + 65)
		} else if 97 <= r && r <= 122 {
			return rune(((a*(int(r)-97+b)+c)%26+26)%26 + 97)
		}
		return r
	}
}

// returns multiplicative inverse mod 26
func affineModInverse(a int) int {
	for x := 1; x <= 25; x++ {
		if (a*x)%26 == 1 {
			return x
		}
	}
	return 0
	//import "math/big"
	//return int(new(big.Int).ModInverse(big.NewInt(int64(a)), big.NewInt(26)).Int64())
}
