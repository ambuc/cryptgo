package main

//import "fmt"
import "strings"

type atbash struct{
  input string
}

func (a atbash) encrypt() (string, error) {
  return strings.Map(flip, a.input), nil
}

func (a atbash) decrypt() (string, error) {
  return strings.Map(flip, a.input), nil
}
