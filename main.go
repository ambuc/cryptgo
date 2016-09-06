package main

import "fmt"
import "io/ioutil"
import "strings"
import "github.com/pborman/getopt" //https://godoc.org/github.com/pborman/getopt
import "errors"
import "os"
import "text/tabwriter"

type cipher interface {
  encrypt() (string, error)
  decrypt() (string, error)
}

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
    return false, errors.New("No input supplied. \n       Try `--input <string>` or `--read <file>`.\n                `-i <string>`        `-r <file>`")
  }
  if(!w.encrypting && !w.decrypting){
    return false, errors.New("Neither encrypting nor decrypting. \n       Try `--encrypt` or `--decrypt`.\n            `-e`           `-d`")
  }
  if(w.encrypting && w.decrypting){
    return false, errors.New("Both encrypting and decrypting.")
  }
  if(!w.existsCipher){
    return false, errors.New("No cipher defined. \n       Try `--cipher (caesar|atbash|rot13)`")
  }

  return true, nil
}

func (w World) print() {
  p := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)

  var status, inputName, outputName string
  if w.encrypting {
    status = "encrypting"; inputName = "Plaintext";  outputName = "Ciphertext"
  } else {
    status = "decrypting"; inputName = "Ciphertext"; outputName = "Plaintext"
  }

  fmt.Println("")
    fmt.Fprintln(p, "Cipher \t", w.cipher)
    fmt.Fprintln(p,  "Status \t", status)
    if w.decrypting && w.hint != "" { fmt.Fprintln(p, "Hint \t", w.hint) }
    if (w.existsInputPath) { fmt.Fprintln(p, "Input path \t", w.inputPath) }
    if (w.existsOutputPath) { fmt.Fprintln(p, "Output path \t", w.outputPath) }
    fmt.Fprintln(p, inputName, "\t", shorten(strings.TrimSpace(w.input)))
    fmt.Fprintln(p, outputName, "\t", shorten(strings.TrimSpace(w.output)))
    if (w.existsOutputPath) { fmt.Fprintln(p, "\t Printed to", w.outputPath) }
  p.Flush()
  if (!w.existsOutputPath) { fmt.Println("\n", w.output) }
  fmt.Println("")
}

func (w World) process() (string, error) {
  var c cipher

  switch w.cipher {
    case "caesar" : c = caesar{input: w.input, n: w.n, hint: w.hint}
    case "rot13"  : c =  rot13{input: w.input}
    case "atbash" : c = atbash{input: w.input}
    default:
      return "", errors.New("No cipher defined. Try --cipher caesar")
  }

  if (w.encrypting) {
    return c.encrypt()
  } else if (w.decrypting) {
    return c.decrypt()
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
  world.inputPath        = *inputPath
  world.existsInputPath  = (*inputPath != "")
  world.outputPath       = *outputPath
  world.existsOutputPath = (*outputPath != "")
  world.input            = *input
  world.existsInput      = (*input != "")
  world.hint             = *hintPtr
  world.n                = *nPtr
  world.args             = getopt.Args()
  world.cipher           = *cipherPtr
  world.existsCipher     = (*cipherPtr != "")

  _, err := world.check()
  check(err)

  if(world.existsInputPath){
    inputTextBytes, inputTextErr := ioutil.ReadFile(world.inputPath)
    check(inputTextErr)
    world.input = strings.TrimSpace(string(inputTextBytes))
  }

  world.output, err = world.process()
  check(err)

  if (world.existsOutputPath) {
    err := ioutil.WriteFile(world.outputPath, []byte(world.output), 0644)
    check(err)
  }

  if (!*quietFlag){
    world.print()
  } else {
    fmt.Println(world.output)
  }


}
