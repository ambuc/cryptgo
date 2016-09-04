# cryptgo
##cryptography package written in go

**Tenative flagset**:
```
 -i --input  <inputFile.txt>
 -o --output <outputFile.txt>
 -e --encode
 -d --decode
```
`-o` is optional -- it should print to terminal by default.
##Caesar Cipher
```
./crypt -e caesar +5 -i plain.txt -o cipher.txt
./crypt -e caesar +5 -i plain.txt -o cipher.txt
./crypt -d caesar    -i plain.txt -o cipher.txt
```
decoding could be performed 
 - with `--print-all` (which prints all 26)
 - or with default, `--dict-attack`
```
./crypt -d caesar --print-all -i plain.txt
./crypt -d caesar --dict-attack -i plain.txt
```
