package main

import "fmt"
import "io/ioutil"
import "strings"
import "github.com/pborman/getopt"
import "errors"

func check(e error) {
  if e != nil {
    panic(e)
  } 
}

type settings struct {
  encrypting bool
  decrypting bool
  existsInputPath bool
  existsOutputPath bool
  inputPath string
  outputPath string
  cipher string
  existsCipher bool
  args []string
  hint string
}

func checkWorld(world settings) (bool, error) {
  if(!world.existsInputPath){
    return false, errors.New("No input supplied. Try `... --inputpath /path/to/input.txt`")
  }
  if(!world.encrypting && !world.decrypting){
    return false, errors.New("Neither encrypting nor decrypting. Try --encrypt or --decrypt")
  }
  if(world.encrypting && world.decrypting){
    return false, errors.New("Both encrypting and decrypting. Try --encrypt or --decrypt")
  }
  if(!world.existsCipher){
    return false, errors.New("No cipher defined. Try `... --cipher caesar`")
  }

  return true, nil
}


func printWorld(world settings, inputText string, outputText string) {
  fmt.Println("")
  fmt.Println("    cipher ::", world.cipher)

  if world.encrypting {
    fmt.Println("    status :: encrypting")
  } else if world.decrypting {
    fmt.Println("    status :: decrypting")
    if (world.hint != "") {
      fmt.Println("      hint ::", world.hint)
    }
  } else {
    fmt.Println("    status :: neither encrypting nor decryptinng")
  }

  fmt.Println(" inputPath ::", world.inputPath)
  if (world.existsOutputPath) {
    fmt.Println("outputPath ::", world.outputPath)
  }

  fmt.Println(" inputText :: <", strings.TrimSpace(inputText), ">")

  if (world.existsOutputPath) {
    fmt.Println("outputText :: printed to", world.outputPath)
  } else {
    fmt.Println("outputText :: <", outputText, ">")
  }

  fmt.Println("")
}

func process(inputText string, world settings) (string, error) {
  switch world.cipher {
  case "caesar":
    if (world.encrypting) {
      return caesarEncrypt(inputText, world.args)
    } else if (world.decrypting) {
      return caesarDecrypt(inputText, world.hint)
    }
  default:
    return "", errors.New("No cipher defined. Try --cipher caesar")
  }
  return "", errors.New("No cipher defined. Try --cipher caesar")
}

func main() {

  encryptingFlag := getopt.BoolLong("encrypt", 'e', "encrypting?")
  decryptingFlag := getopt.BoolLong("decrypt", 'd', "decrypting?")
  quietFlag      := getopt.BoolLong("quiet", 'q', "quiet?")
  inputPath      := getopt.StringLong("inputpath", 'i', "", "path to input file")
  outputPath     := getopt.StringLong("outputpath",'o', "", "path to output file")
  cipherPtr      := getopt.StringLong("cipher", 'c', "", "which cipher to use")
  hintPtr        := getopt.StringLong("hint", 'h', "", "hint for the decrypter")

  getopt.Parse()

  world                 := settings{}
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

  worldOk, err := checkWorld(world)
  check(err)
  if(!worldOk){
    return
  }

  inputTextBytes, inputTextErr := ioutil.ReadFile(*inputPath)
  check(inputTextErr)
  inputText := strings.TrimSpace(string(inputTextBytes))

  outputText, err := process(inputText, world)
  check(err)

  if (!*quietFlag){
    printWorld(world, inputText, outputText)
  }

  if (world.existsOutputPath) {
    err := ioutil.WriteFile(*outputPath, []byte(outputText), 0644)
    check(err)
  }

  //fmt.Println("args:", getopt.Args())
  //fmt.Println(reflect.TypeOf(getopt.Args()))

}
