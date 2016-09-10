package main

import "errors"
import "fmt"
import "github.com/pborman/getopt" //https://godoc.org/github.com/pborman/getopt
import "io/ioutil"
import "os"
import "strings"
import "text/tabwriter"

type cipher interface {
	encrypt() (string, error)
	decrypt() (string, error)
}

type World struct {
	encrypting bool
	decrypting bool

	inputPath  string
	outputPath string
	input      string
	output     string
	cipher     string

	hint string
	key  string

	n int
	a int
	b int

	args []string
}

func (w World) process() (string, error) {
	var c cipher

	switch w.cipher {
	case "caesar":
		c = caesar{input: w.input, n: w.n, hint: w.hint}
	case "rot13":
		c = rot13{input: w.input}
	case "atbash":
		c = atbash{input: w.input}
	case "affine":
		c = affine{input: w.input, hint: w.hint, a: w.a, b: w.b}
	case "substitution":
		c = substitution{input: w.input, hint: w.hint, key: w.key}
	default:
		return "", errors.New("No cipher defined. Try --cipher caesar")
	}

	if w.encrypting {
		return c.encrypt()
	} else if w.decrypting {
		return c.decrypt()
	}

	return "", errors.New("No cipher defined. Try --cipher caesar")
}

func main() {

	encryptingFlag := getopt.BoolLong("encrypt", 'e', "Boolean, true if encrypting the input.")
	decryptingFlag := getopt.BoolLong("decrypt", 'd', "Boolean, true if decrypting t he input.")
	inputPath := getopt.StringLong("read", 'r', "", "Path to input file.")
	outputPath := getopt.StringLong("write", 'w', "", "Path to output file. (optional)")
	input := getopt.StringLong("input", 'i', "", "Input as a string.")
	cipherPtr := getopt.StringLong("cipher", 'c', "", "Name of encryption/decryption method used.")
	hintPtr := getopt.StringLong("hint", 'h', "", "Hint for the decrypter, varies across ciphers. (optional)")
	keyPtr := getopt.StringLong("key", 'k', "", "Some ciphers require a passphrase or key.")
	nPtr := getopt.IntLong("num", 'n', 0, "Some ciphers require a shift by <n> characters.")
	aPtr := getopt.IntLong("a", 'a', 0, "Some ciphers require keys.")
	bPtr := getopt.IntLong("b", 'b', 0, "Some ciphers require keys.")
	quietFlag := getopt.BoolLong("quiet", 'q', "Boolean, true if suppressing verbose output.")

	getopt.Parse()

	world := World{}
	world.encrypting = *encryptingFlag
	world.decrypting = *decryptingFlag
	world.inputPath = *inputPath
	world.outputPath = *outputPath
	world.input = *input
	world.hint = *hintPtr
	world.key = *keyPtr
	world.n = *nPtr
	world.a = *aPtr
	world.b = *bPtr
	world.args = getopt.Args()
	world.cipher = *cipherPtr

	if world.inputPath == "" && world.input == "" {
		panic(errors.New("No input supplied. \n       Try `--input <string>` or `--read <file>`.\n                `-i <string>`        `-r <file>`"))
	}
	if !world.encrypting && !world.decrypting {
		panic(errors.New("Neither encrypting nor decrypting. \n       Try `--encrypt` or `--decrypt`.\n            `-e`           `-d`"))
	}
	if world.encrypting && world.decrypting {
		panic(errors.New("Both encrypting and decrypting. \n       Try `--encrypt` or `--decrypt`.\n            `-e`           `-d`"))
	}
	if world.cipher == "" {
		panic(errors.New("No cipher defined. \n       Try `--cipher (caesar|atbash|rot13)`"))
	}

	if world.inputPath != "" {
		inputTextBytes, inputTextErr := ioutil.ReadFile(world.inputPath)
		check(inputTextErr)
		world.input = strings.TrimSpace(string(inputTextBytes))
	}

	var err error
	world.output, err = world.process()
	check(err)

	if world.outputPath != "" {
		err = ioutil.WriteFile(world.outputPath, []byte(world.output), 0644)
		check(err)
	}

	if !*quietFlag {

		var status, inputName, outputName string
		if world.encrypting {
			status = "encrypting"
			inputName = "Plaintext"
			outputName = "Ciphertext"
		} else {
			status = "decrypting"
			inputName = "Ciphertext"
			outputName = "Plaintext"
		}

		fmt.Println("")

		p := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)
		fmt.Fprintln(p, "Cipher \t", world.cipher)
		fmt.Fprintln(p, "Status \t", status)
		if world.decrypting && world.hint != "" {
			fmt.Fprintln(p, "Hint \t", world.hint)
		}
		if world.key != "" {
			fmt.Fprintln(p, "Key \t", world.key)
		}
		if world.inputPath != "" {
			fmt.Fprintln(p, "Input path \t", world.inputPath)
		}
		if world.outputPath != "" {
			fmt.Fprintln(p, "Output path \t", world.outputPath)
		}
		fmt.Fprintln(p, inputName, "\t", shorten(strings.TrimSpace(world.input)))
		fmt.Fprintln(p, outputName, "\t", shorten(strings.TrimSpace(world.output)))
		if world.outputPath != "" {
			fmt.Fprintln(p, "\t Printed to", world.outputPath)
		}
		p.Flush()

		fmt.Println("")
	} else {
		fmt.Println(world.output)
	}

}
