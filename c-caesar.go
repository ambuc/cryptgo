package main

import "strconv"
import "errors"

func caesarEncrypt(inputText string, args []string) (string, error) {
  n := 0; var err error
  if (len(args) != 0){
    n, err = strconv.Atoi(args[0])
    check(err)
  }
  if (n==0) {
    return shiftWord(inputText, n), errors.New("no shift found. try `--cipher caesar 5`")
  }
  return shiftWord(inputText, n), err
}

func caesarDecrypt(inputText string, hint string) (string, error) {
  switch hint {

    case "brute-force":
      result := "\n"
      i := 0
      for (i < 26) {
        result = result + shiftWord(inputText, i) + "\n"
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
