package main

import "fmt"
import "io/ioutil"
import "strings"
import "github.com/pborman/getopt"


func check(e error) {
  if e != nil {
    panic(e)
  } 
}

type sett struct {
  encrypting bool
  decrypting bool
  existsInputPath bool
  existsOutputPath bool
  inputPath string
  outputPath string
  cipher string
}

func printStatus(settings sett, inputText string, outputText string) {
  fmt.Println("")
  fmt.Println("    cipher ::", settings.cipher)

  //print coding status
  if settings.encrypting {
    fmt.Println("    status :: encrypting")
  } else if settings.decrypting {
    fmt.Println("    status :: decrypting")
  } else {
    fmt.Println("    status :: neither encrypting nor decryptinng")
  }
  
  //print in/output paths
  fmt.Println(" inputPath ::", settings.inputPath)
  if (settings.existsOutputPath) {
    fmt.Println("outputPath ::", settings.outputPath)
  }

  fmt.Println(" inputText :: <", strings.TrimSpace(inputText), ">")

  if (settings.existsOutputPath) {
    fmt.Println("outputText :: printed to", settings.outputPath)
  } else {
    fmt.Println("outputText ::", outputText)
  }

  fmt.Println("")
}


func main() {

  encryptingFlag := getopt.BoolLong("encrypt", 'e', "encrypting?")
  decryptingFlag := getopt.BoolLong("decrypt", 'd', "decrypting?")
  quietFlag      := getopt.BoolLong("quiet", 'q', "quiet?")
  inputPath      := getopt.StringLong("inputpath", 'i', "path to input file")
  outputPath     := getopt.StringLong("outputpath",'o', "path to output file")
  cipherPtr      := getopt.StringLong("cipher", 'c', "which cipher to use")

  getopt.Parse()

  settings := sett{}
  settings.encrypting = *encryptingFlag
  settings.decrypting = *decryptingFlag
  settings.cipher = *cipherPtr
  settings.inputPath = *inputPath
  settings.outputPath = *outputPath
  settings.existsInputPath = (*inputPath != "")
  settings.existsOutputPath = (*outputPath != "")

  inputTextBytes, inputTextErr := ioutil.ReadFile(*inputPath)
  check(inputTextErr)
  inputText := string(inputTextBytes)

  outputText := "whatever, man"

  if (settings.existsOutputPath) {
    err := ioutil.WriteFile(*outputPath, []byte(outputText), 0644)
    check(err)
  }

  if (!*quietFlag){
    printStatus(settings, inputText, outputText)
  }

}
