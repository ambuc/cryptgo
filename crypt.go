package main

import "fmt"
import "io/ioutil"
import "strings"
import "github.com/pborman/getopt"
import "strconv"
//import "reflect"

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
  args []string
  //caesar flags
  printAll bool
}

func printStatus(world settings, inputText string, outputText string) {
  fmt.Println("")
  fmt.Println("    cipher ::", world.cipher)

  if world.encrypting {
    fmt.Println("    status :: encrypting")
  } else if world.decrypting {
    fmt.Println("    status :: decrypting")
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

func shift(r rune, shift int) rune {
  //fmt.Println(string(r), r, int(r)+shift, string(int(r)+shift))
  if( 65<=r && r<=90 ) {
    tmp := int(r) - 65
    tmp = tmp + shift
    tmp = tmp % 26
    return rune(tmp + 65)
  }
  if( 97<=r && r<=122 ) {
    tmp := int(r) - 97
    tmp = tmp + shift
    tmp = tmp % 26
    return rune(tmp + 97)
  }
  return r
}

func caesarEncrypt(inputText string, n int) string {
  return strings.Map(func(r rune) rune {
    return shift(r, n)
  }, inputText)
}

func caesarDecrypt(inputText string) string {
  i := 0
  for (i < 26) {
    fmt.Println(caesarEncrypt(inputText, i))
    i = i + 1
  }
  return ""
}

func process(inputText string, world settings) string {
  if (world.cipher == "caesar"){
    if (world.encrypting) {
      n, err := strconv.Atoi(world.args[0])
      check(err)
      return caesarEncrypt(inputText, n)
    } else if (world.decrypting) {
      return caesarDecrypt(inputText)
    }
  } else {
    return "no such cipher"
  }
  return ""
}

func main() {

  encryptingFlag := getopt.BoolLong("encrypt", 'e', "encrypting?")
  decryptingFlag := getopt.BoolLong("decrypt", 'd', "decrypting?")
  quietFlag      := getopt.BoolLong("quiet", 'q', "quiet?")
  inputPath      := getopt.StringLong("inputpath", 'i', "", "path to input file")
  outputPath     := getopt.StringLong("outputpath",'o', "", "path to output file")
  cipherPtr      := getopt.StringLong("cipher", 'c', "which cipher to use")

  getopt.Parse()

  world                 := settings{}
  world.encrypting       = *encryptingFlag
  world.decrypting       = *decryptingFlag
  world.cipher           = *cipherPtr
  world.inputPath        = *inputPath
  world.outputPath       = *outputPath
  world.existsInputPath  = (*inputPath != "")
  world.existsOutputPath = (*outputPath != "")
  world.args             = getopt.Args()


  inputTextBytes, inputTextErr := ioutil.ReadFile(*inputPath)
  check(inputTextErr)
  inputText := strings.TrimSpace(string(inputTextBytes))

  outputText := process(inputText, world)

  if (!*quietFlag){
    printStatus(world, inputText, outputText)
  }

  if (world.existsOutputPath) {
    err := ioutil.WriteFile(*outputPath, []byte(outputText), 0644)
    check(err)
  }

  //fmt.Println("args:", getopt.Args())
  //fmt.Println(reflect.TypeOf(getopt.Args()))

}
