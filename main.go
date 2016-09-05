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

        input string
  existsInput bool
       output string
 existsOutput bool

        cipher string
  existsCipher bool

  hint string
  n int
  args []string
}

func (w World) check() (bool, error) {
  if(!w.existsInputPath && !w.existsInput){
    return false, errors.New("No input supplied. Try `--inputpath <path>` or `--input <input>`")
  }
  if(!w.encrypting && !w.decrypting){
    return false, errors.New("Neither encrypting nor decrypting. Try --encrypt or --decrypt")
  }
  if(w.encrypting && w.decrypting){
    return false, errors.New("Both encrypting and decrypting. Try --encrypt or --decrypt")
  }
  if(!w.existsCipher){
    return false, errors.New("No cipher defined. Try `--cipher caesar`")
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

  if (w.existsInputPath) {
    fmt.Fprintln(p, "Input path \t", w.inputPath)
  }
  if (w.existsOutputPath) {
    fmt.Fprintln(p, "Output path \t", w.outputPath)
  }

  if (w.encrypting) {
    fmt.Fprintln(p, " Plaintext \t", shorten(strings.TrimSpace(w.input)))
  } else {
    fmt.Fprintln(p, " Ciphertext \t", shorten(strings.TrimSpace(w.input)))
  }

  
  if (w.existsOutputPath) {
    if (w.encrypting) {
      fmt.Fprintln(p,  "Ciphertext \t", shorten(strings.TrimSpace(w.output)))
    } else { 
      fmt.Fprintln(p,  "Plaintext \t", shorten(strings.TrimSpace(w.output)))
    }
    fmt.Fprintln(p, "\t Printed to", w.outputPath)
  }

  p.Flush()
  
  if (!w.existsOutputPath) { 
    fmt.Println("")
    fmt.Println(w.output) 
  }
  fmt.Println("")
}

func (w World) process() (string, error) {
  switch w.cipher {
    case "caesar":
      if (w.encrypting) {
        return caesarEncrypt(w.input, w.n)
      } else if (w.decrypting) {
        return caesarDecrypt(w.input, w.hint)
      }
    default:
      return "", errors.New("No cipher defined. Try --cipher caesar")
  }
  return "", errors.New("No cipher defined. Try --cipher caesar")
}

func main() {

  encryptingFlag := getopt.BoolLong("encrypt", 'e', "Boolean, true if encrypting the input.")
  decryptingFlag := getopt.BoolLong("decrypt", 'd', "Boolean, true if decrypting t he input.")
  inputPath      := getopt.StringLong("read", 'r', "", "Path to input file.")
  outputPath     := getopt.StringLong("write",'w', "", "Path to output file. (optional)")
  input          := getopt.StringLong("input", 'i', "", "Input as a string.")
  cipherPtr      := getopt.StringLong("cipher", 'c', "", "Name of encryption/decryption method used.")
  hintPtr        := getopt.StringLong("hint", 'h', "", "Hint for the decrypter, varies across ciphers. (optional)")
  nPtr           := getopt.IntLong("num", 'n', 0, "Some ciphers require a shift by <n> characters.")

  quietFlag      := getopt.BoolLong("quiet", 'q', "Boolean, true if suppressing verbose output.")

  getopt.Parse()

  world                 := World{}
  world.encrypting       = *encryptingFlag
  world.decrypting       = *decryptingFlag
  world.cipher           = *cipherPtr
  world.existsCipher     = (*cipherPtr != "")
  world.inputPath        = *inputPath
  world.outputPath       = *outputPath
  world.input            = *input
  world.existsInputPath  = (*inputPath != "")
  world.existsOutputPath = (*outputPath != "")
  world.existsInput      = (*input != "")
  world.hint             = *hintPtr
  world.n                = *nPtr
  world.args             = getopt.Args()

  worldOk, err := world.check()
  check(err)
  if(!worldOk){ return }

  if(world.existsInputPath){
    inputTextBytes, inputTextErr := ioutil.ReadFile(*inputPath)
    check(inputTextErr)
    world.input = strings.TrimSpace(string(inputTextBytes))
  }
  

  world.output, err = world.process()
  check(err)

  if (!*quietFlag){
    world.print()
  } else {
    fmt.Println(world.output)
  }

  if (world.existsOutputPath) {
    err := ioutil.WriteFile(*outputPath, []byte(world.output), 0644)
    check(err)
  }

}
