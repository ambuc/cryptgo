# cryptgo
##cryptography package written in go


##Usage:

```
 ./cryptgo -r input.txt (-w output.txt) -e -c caesar -n 5
                                        -d -c caesar -h analyze
 ./cryptgo -i plaintext  -e -c caesar -n 5
 ./cryptgo -i ciphertext -d -c caesar -h analyze

Options:
  -e --encrypt             Boolean, true if encrypting the input.
  -d --decrypt             Boolean, true if decrypting the input.
  -r --read=<path>         Path to input file.
  -w --write=<path>        Output is printed to the shell by default, but can be directed. (optional)
                           into an output file.
  -i --input=<string>      Input as a string.
  -h --hint=<string>       Hint for the decrypter, varies across ciphers. (optional)
  -c --cipher=<string>     Name of encryption/decryption method used.
  -n --num=<num>           Some ciphers require a shift by <n> characters
  -q --quiet               Boolean, true if suppressing verbose output.
```
Short forms are available for all flags, but we use long flags in the documentation below.

## Supported ciphers
 - [x] [Caesar cipher](#caesar-cipher) `-c caesar`
 - [x] [ROT13](#rot13-cipher) `-c rot13`
 - [ ] Substitution cipher
 - [ ] Atbash cipher
 - [ ] Affine cipher
 - [ ] Rail Fence cipher
 - [ ] Route cipher
 - [ ] Vignere cipher
 - [ ] Playfair cipher
 - [ ] Hill cipher

###Caesar cipher
####Encryption
```
 ./cryptgo --input plaintext --encrypt --cipher caesar --num 5
           --read input.txt
```
####Decryption
```
 ./cryptgo --read input.txt --decrypt --cipher caesar --hint brute-force
                                                      --hint analyze
                                                      --hint analyze-verbose
```
 - The _brute force_ method simply prints all 26 options, unranked.
 - The _analyze_ method runs a simple frequency analysis against the sample (assuming english), and returns the best fit. 
 - The _analyze-verbose_ method returns the best fit to output, but prints all 26 ranked options to stdout.

###ROT13 cipher
```
 ./cryptgo --read input.txt --encrypt --cipher rot13
                            --decrypt
```
##Installation
```
git clone https://github.com/ambuc/cryptgo.git
cd cryptgo
go build -o crypt *.go
./cryptgo ...
```

##Dependencies
Requires [`getopt`](https://godoc.org/github.com/pborman/getopt).
```
go get github.com/pborman/getopt
```

