package main

import "flag"
import "fmt"
import "io/ioutil"
import "strings"

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {

  //read encode/decode pointer flag
  isEncodePtr := flag.Bool("e", false, "encode? false by default")
  isDecodePtr := flag.Bool("d", false, "decode? true by default")

  //read input/outut path strings
  var inputPath string
  flag.StringVar(&inputPath, "i", "", "path to input file")
  var outputPath string
  flag.StringVar(&outputPath, "o", "", "path to output file")

  flag.Parse()

  //print coding status
  if (*isEncodePtr == true) {
    fmt.Println("    status :: encoding")
  } else if *isDecodePtr {
    fmt.Println("    status :: decoding")
  } else {
    fmt.Println("    status :: neither encoding nor decoding")
  }

  //print inputPath and get inputText
  fmt.Println(" inputPath ::", inputPath)
  if (outputPath != "") {
    fmt.Println("outputPath ::", outputPath)
  }
  inputTextBytes, inputTextErr := ioutil.ReadFile(inputPath)
  check(inputTextErr)
  inputText := string(inputTextBytes)

  fmt.Println(" inputText ::", strings.TrimSpace(inputText))

  outputText := "whatever, man"

  fmt.Println("outputText ::", outputText)

  if (outputPath != "") {
    fmt.Println("output printed to", outputPath)
    err := ioutil.WriteFile(outputPath, []byte(outputText), 0644)
    check(err)
  } else {
    fmt.Println("Output path not defined, printing output:")
    fmt.Println(outputText)
  }

}
