package main

import "fmt"
import "io/ioutil"
import "strings"
import "github.com/pborman/getopt" //https://godoc.org/github.com/pborman/getopt
import "errors"
import "os"
import "text/tabwriter"


type World struct {
  encrypting bool
  decrypting bool

        inputPath string
  existsInputPath bool
        outputPath string
  existsOutputPath bool

  inputText string
  outputText string

        cipher string
  existsCipher bool

  hint string
  args []string
}

func (w World) check() (bool, error) {
  if(!w.existsInputPath){
    return false, errors.New("No input supplied. Try `... --inputpath /path/to/input.txt`")
  }
  if(!w.encrypting && !w.decrypting){
    return false, errors.New("Neither encrypting nor decrypting. Try --encrypt or --decrypt")
  }
  if(w.encrypting && w.decrypting){
    return false, errors.New("Both encrypting and decrypting. Try --encrypt or --decrypt")
  }
  if(!w.existsCipher){
    return false, errors.New("No cipher defined. Try `... --cipher caesar`")
  }

  return true, nil
}

func (w World) print() {
  p := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)
  fmt.Println("")

  fmt.Fprintln(p, "Cipher \t", w.cipher)

  if w.encrypting {
    fmt.Fprintln(p,  "Status \t encrypting")
  } else if w.decrypting {
    fmt.Fprintln(p, "Status \t decrypting")
    if (w.hint != "") {
      fmt.Fprintln(p, "Hint \t", w.hint)
    }
  } else {
    fmt.Fprintln(p, "Status \t neither encrypting nor decrypting")
  }

  fmt.Fprintln(p, "Input path \t", w.inputPath)
  if (w.existsOutputPath) {
    fmt.Fprintln(p, "Output path \t", w.outputPath)
  }

  if (w.encrypting) {
    fmt.Fprintln(p, " Plaintext \t", shorten(strings.TrimSpace(w.inputText)))
  } else {
    fmt.Fprintln(p, " Ciphertext \t", shorten(strings.TrimSpace(w.inputText)))
  }

  
  if (w.existsOutputPath) {
    if (w.encrypting) {
      fmt.Fprintln(p,  "Ciphertext \t printed to", w.outputPath)
    } else { 
      fmt.Fprintln(p,  "Plaintext \t printed to", w.outputPath)
    }
  }

  p.Flush()
  
  fmt.Println("")
  if (!w.existsOutputPath) { fmt.Println(w.outputText) }
  fmt.Println("")
}

func (w World) process() (string, error) {
  switch w.cipher {
    case "caesar":
      if (w.encrypting) {
        return caesarEncrypt(w.inputText, w.args)
      } else if (w.decrypting) {
        return caesarDecrypt(w.inputText, w.hint)
      }
    default:
      return "", errors.New("No cipher defined. Try --cipher caesar")
  }
  return "", errors.New("No cipher defined. Try --cipher caesar")
}

func main() {

  encryptingFlag := getopt.BoolLong("encrypt", 'e', "encrypting?")
  decryptingFlag := getopt.BoolLong("decrypt", 'd', "decrypting?")
  inputPath      := getopt.StringLong("inputpath", 'i', "", "path to input file")
  outputPath     := getopt.StringLong("outputpath",'o', "", "path to output file")
  cipherPtr      := getopt.StringLong("cipher", 'c', "", "which cipher to use")
  hintPtr        := getopt.StringLong("hint", 'h', "", "hint for the decrypter")

  quietFlag      := getopt.BoolLong("quiet", 'q', "quiet?")

  getopt.Parse()

  world                 := World{}
  world.encrypting       = *encryptingFlag
  world.decrypting       = *decryptingFlag
  world.cipher           = *cipherPtr
  world.existsCipher     = (*cipherPtr != "")
  world.inputPath        = *inputPath
  world.outputPath       = *outputPath
  world.existsInputPath  = (*inputPath != "")
  world.existsOutputPath = (*outputPath != "")
  world.hint             = *hintPtr
  world.args             = getopt.Args()

  worldOk, err := world.check()
  check(err)
  if(!worldOk){ return }

  inputTextBytes, inputTextErr := ioutil.ReadFile(*inputPath)
  check(inputTextErr)
  world.inputText = strings.TrimSpace(string(inputTextBytes))

  world.outputText, err = world.process()
  check(err)

  if (!*quietFlag){
    world.print()
  } else {
    fmt.Println(world.outputText)
  }

  if (world.existsOutputPath) {
    err := ioutil.WriteFile(*outputPath, []byte(world.outputText), 0644)
    check(err)
  }

}
