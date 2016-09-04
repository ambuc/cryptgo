package main

import "flag"
import "fmt"
import "io/ioutil"

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
  inputText, inputTextErr := ioutil.ReadFile(inputPath)
  check(inputTextErr)

  fmt.Print(string(inputText))
}
