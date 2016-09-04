package main

import "flag"
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
  method string
}

func printStatus(settings sett, inputText string, outputText string) {
  fmt.Println("")
  fmt.Println("    method ::", settings.method)

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

  //read encode/decode pointer flag
  encryptingPtr := flag.Bool("e", false, "encode? false by default")
  decryptingPtr := flag.Bool("d", false, "decode? true by default")
  quietPtr      := flag.Bool("q", false, "quiet? false by default")

  //read input/outut path strings
  var  inputPath string; flag.StringVar( &inputPath, "i", "", "path to input file" )
  var outputPath string; flag.StringVar(&outputPath, "o", "", "path to output file")

  flag.Parse()

  settings := sett{}
  settings.encrypting = *encryptingPtr
  settings.decrypting = *decryptingPtr
  settings.method = flag.Args()[0]
  settings.inputPath = inputPath
  settings.outputPath = outputPath
  settings.existsInputPath = (inputPath != "")
  settings.existsOutputPath = (outputPath != "")


  //print inputPath and get inputText
  inputTextBytes, inputTextErr := ioutil.ReadFile(inputPath)
  check(inputTextErr)
  inputText := string(inputTextBytes)

  outputText := "whatever, man"

  if (settings.existsOutputPath) {
    err := ioutil.WriteFile(outputPath, []byte(outputText), 0644)
    check(err)
  }

  if (!*quietPtr){
    printStatus(settings, inputText, outputText)
  }

}
