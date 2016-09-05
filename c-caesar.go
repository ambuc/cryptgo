package main

import "strings"
import "strconv"
import "errors"

func caesarShift(inputText string, n int) string {
  return strings.Map( func (r rune) rune { return shift(r, n) }, inputText)
}

func caesarEncrypt(inputText string, args []string) (string, error) {
  n := 0; var err error
  if (len(args) != 0){
    n, err = strconv.Atoi(args[0])
    check(err)
  }
  if (n==0){
    return caesarShift(inputText, n), errors.New("no shift found. try `--cipher caesar 5`")
  }
  return caesarShift(inputText, n), err
}

func caesarDecrypt(inputText string, hint string) (string, error) {
  switch hint {
  case "brute-force":
    result := "\n"
    i := 0
    for (i < 26) {
      result = result + caesarShift(inputText, i) + "\n"
      i = i + 1
    }
    return result, nil
  case "analyze":
    return "working on it", nil
  default:
    return "", errors.New("no hint given. specify --hint brute-force or --hint analyze")
  }
  return "", nil
}
