package main

import "strings"

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

