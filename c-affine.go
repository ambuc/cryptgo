package main

import "fmt"
import "errors"
import "math/big"
import "strings"

type affine struct {
	hint  string
	input string
	a     int
	b     int
}

func (a affine) encrypt() (string, error) {
	if a.a <= 0 {
		return "", errors.New("Insecure encryption: value <a> must be greater than zero")
	}
	if a.b <= 0 {
		return "", errors.New("Insecure encryption: value <b> must be greater than zero")
	}
	if !GCDIterative(26, a.a) {
		return "", errors.New("Insecure encryption: value <a> not coprime to 26")
	}
	return strings.Map(affineShift(a.a, 0, a.b), a.input), nil
}

func (a affine) decrypt() (string, error) {
	switch a.hint {
	case "known":
		if a.a <= 0 {
			return "", errors.New("Insecure encryption: value <a> must be greater than zero")
		}
		if a.b <= 0 {
			return "", errors.New("Insecure encryption: value <b> must be greater than zero")
		}
		if !GCDIterative(26, a.a) {
			return "", errors.New("Unreal decryption: <a> not coprime to 26")
		}
		j := int(new(big.Int).ModInverse(big.NewInt(int64(a.a)), big.NewInt(26)).Int64())
		return strings.Map(affineShift(j, -a.b, 0), a.input), nil

	case "analyze":
		fmt.Print(a.input)
		return a.input, nil
	case "analyze-verbose":
		return a.input, nil
	}
	return "", errors.New("no hint given. specify `--hint known` or `--hint analyze` or `--hint analyize-verbose`")
	return a.input, nil
}
