# cryptgo
##cryptography package written in go


**Tenative flagset**:
```
 -i --input  <inputFile.txt>
 -o --output <outputFile.txt>
 -e --encrypt
 -d --decrypt
 -q --quiet
```
`-o` is optional -- it should print to terminal by default.
##Caesar Cipher
```
./crypt --input-file plain.txt --encrypt --cipher caesar 5
./crypt -i plain.txt -e -c caesar 5
```
decoding could be performed 
 - with `--print-all` (which prints all 26)
 - or with default, `--dict-attack`
```
./crypt -i plain.txt --decrypt --cipher caesar --print-all
./crypt -i plain.txt -d -c caesar --dict-attack
```

##Dependencies
Requires [`getopt`](https://godoc.org/github.com/pborman/getopt).
```
go get github.com/pborman/getopt
```

