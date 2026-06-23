package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Empty string by default
	var str string

	// String with escape sequences
	var hello string = "Hello\n\t"

	// Raw string without escape sequences
	var world string = `World\n\t`

	// UTF-8 out of the box
	var helloWorld = "Hello, World!"

	// Japanese language
	hi := "こんにちは"

	// Single quotes are used for bytes and runes
	var rawBinary byte = '\x27'

	// rune is an alias for int32 and is used for Unicode characters
	var someChinese rune = '中'

	fmt.Println(str, hello, world, hi, rawBinary, someChinese)

	// String concatenation
	andGoodMorning := helloWorld + " and Good Morning!"

	fmt.Println(andGoodMorning)

	// Strings are immutable
	// helloWorld[0] = 'H' // This is not allowed

	// Getting the length of the string
	byteLen := len(helloWorld)                    // length in bytes
	symbols := utf8.RuneCountInString(helloWorld) // length in runes

	fmt.Println(byteLen, symbols)

	// Getting a substring by bytes, not by characters
	helloPart := helloWorld[:5] // "Hello"
	firstByte := helloWorld[0]  // byte: 72, not "H"

	fmt.Println(helloPart, firstByte)

	// Converting string to []byte and back
	byteString := []byte(helloWorld)
	helloWorld = string(byteString)

	fmt.Println(helloWorld)
}
