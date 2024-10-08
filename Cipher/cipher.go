package Cipher

import (
	"fmt"
)

const Flag = false
const ExpectedCRC32 = 0x196d401e
const ExpectedCRC32_2 = 0x7a244a08

const key = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98" //some random numbers here

func Encrypt(input string) (output string) {
	for i := 0; i < len(input); i++ {
		output += fmt.Sprintf("\\x%02x", input[i]^key[i%len(key)])
	}
	return output
}

func Decrypt(input string) (output string) {
	for i := 0; i < len(input); i++ {
		output += string(input[i] ^ key[i%len(key)])
	}
	return output
}
