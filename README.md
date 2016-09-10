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
 - [x] [Atbash cipher](#atbash-cipher) `-c atbash`
 - [ ] Substitution cipher
 - [ ] [Affine cipher](#affine-cipher) `-c affine`
 - [ ] Rail Fence cipher
 - [ ] Route cipher
 - [ ] Vignere cipher
 - [ ] Playfair cipher
 - [ ] Hill cipher

## Hinted encryption / decryption
###[Caesar cipher](https://en.wikipedia.org/wiki/Caesar_cipher)
```
 ./cryptgo --read input.txt --encrypt --cipher caesar --num 5
                            --decrypt --cipher caesar --hint brute-force
                                                      --hint analyze
                                                      --hint analyze-verbose
```
 - The _brute force_ method simply prints all 26 options, unranked.
 - The _analyze_ method runs a simple frequency analysis against the sample (assuming english), and returns the best fit. 
 - The _analyze-verbose_ method returns the best fit to output, but prints all 26 ranked options to stdout.

###[Affine cipher](https://en.wikipedia.org/wiki/Affine_cipher)
  ```
   ./cryptgo --read input.txt --encrypt --cipher affine -a 5 -b 8
                              --decrypt --cipher affine --hint known -a 5 -b 8
                                                        --hint analyze
                                                        --hint analyze-verbose -n 10
  ```
  - The _known_ method assumes prior knowledge of the keys.
  - The _analyze_ method applies [statistical cryptanalysis](http://practicalcryptography.com/cryptanalysis/stochastic-searching/cryptanalysis-affine-cipher/) to determine the best fit.
  - The _analyze-verbose_ method returns the best fit to output, and prints the top `n` best fits.

## Automatic encryption / decryption

###[ROT13 cipher](https://en.wikipedia.org/wiki/ROT13)
  ```
   ./cryptgo --read input.txt --encrypt --cipher rot13
                              --decrypt
  ```
###[Atbash cipher](https://en.wikipedia.org/wiki/Atbash)
  ```
   ./cryptgo --read input.txt --encrypt --cipher atbash
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

