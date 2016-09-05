package main

import "errors"

func caesarEncrypt(inputText string, n int) (string, error) {
  if (n==0) {
    return shiftWord(inputText, n), errors.New("no shift found. try `--cipher caesar -n 5`")
  }
  return shiftWord(inputText, n), nil
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
      return frequencyAnalysis(inputText, false), nil

    case "analyze-verbose":
      return frequencyAnalysis(inputText, true), nil

    default:
      return "", errors.New("no hint given. specify `--hint brute-force` or `--hint analyze` or `--hint analyize-verbose`")
  }
  return "", nil
}
