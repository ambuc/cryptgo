# cryptgo
##cryptography package written in go


##Usage:

```
 ./cryptgo -i=input.txt (-o=output.txt) -e -c=caesar 5
                                        -d
 ./cryptgo --inputpath=input.txt (--outputpath=output.txt) --encrypt --cipher=caesar 5
                                                           --decrypt

Options:
  -e --encrypt             Boolean, true if encrypting the input file.
  -d --decrypt             Boolean, true if decrypting the input file.
  -i --inputpath=<path>    Path to the input file.
  -o --outputpath=<path>   Path to the output file. Output is printed to the shell
                           by default, but can be directed into a file.
  -h --hint=<hint>         Hint for the decrypter, varies across ciphers [optional]
  -c --cipher=<cipher>     Name of encryption/decryption method used.
               caesar <n>  Caesar ciphers require a shift on encryption.
```

##Caesar Cipher
Decryption has three hints available:
```
 ./cryptgo --input input.txt --decrypt --cipher caesar 5 --hint brute-force
 ./cryptgo --input input.txt --decrypt --cipher caesar 5 --hint analyze
 ./cryptgo --input input.txt --decrypt --cipher caesar 5 --hint analyze-verbose
```
 - The _brute force_ method simply prints all 26 options, unranked.
 - The _analyze_ method runs a simple frequency analysis against the sample (assuming english), and returns the best fit. 
 - The _analyze-verbose_ method returns the best fit to output, but prints all 26 ranked options to stdout.

##Installation
```
git clone https://github.com/ambuc/cryptgo.git
cd cryptgo
go build
./cryptgo ...
```

##Dependencies
Requires [`getopt`](https://godoc.org/github.com/pborman/getopt).
```
go get github.com/pborman/getopt
```

