package main

//import "fmt"
import "errors"
//import "math"
import "strings"

type affine struct {
  hint string
  input string
  a int
  b int
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


func (a affine) encrypt() (string, error) {
  coprime := GCDIterative(26, a.a)
  if !coprime {
    return "", errors.New("Insecure encryption: value <a> not coprime to 26")
  }
  if a.b == 0 {
    return "", errors.New("Insecure encryption: value <b> not nonzero")
  }

  
  return strings.Map(func(r rune) rune {
    if( 65<=r && r<=90 ) {
      return rune((a.a*int(r-65)+a.b)%26+65)
    }
    if( 97<=r && r<=122 ) {
      return rune((a.a*int(r-97)+a.b)%26+97)
    }
    return r
  }, a.input), nil
}

func (a affine) decrypt() (string, error) {
  return a.input, nil
}
