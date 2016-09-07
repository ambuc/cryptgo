package main

import "errors"
import "math/big"
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
  switch a.hint{
  case "known":
    coprime := GCDIterative(26, a.a)
    if !coprime {
      return "", errors.New("Unreal decryption: value <a> not coprime to 26")
    }
    i := new(big.Int).ModInverse(big.NewInt(int64(a.a)), big.NewInt(26))
    j := int(i.Int64())
    return strings.Map(func(r rune) rune {
      if( 65<=r && r<=90 ) {
        return rune(((j*(int(r)-65-a.b))%26+26)%26+65)
      }
      if( 97<=r && r<=122 ) {
        return rune(((j*(int(r)-97-a.b))%26+26)%26+97)
      }
      return r
    }, a.input), nil
  case "analyze":
    return a.input, nil
  case "analyze-verbose":
    return a.input, nil
  }
  return "", errors.New("no hint given. specify `--hint known` or `--hint analyze` or `--hint analyize-verbose`")
  return a.input, nil
}
