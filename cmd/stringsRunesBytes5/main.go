package main

import "fmt"

func main() {
	var myString = "résumé" // includees non ascii character é

	// this gets the ith byte character, while range gets the actual character
	// thats why this is one byte even tho é is two bytes
	// it only returns the first half of the character
	var indexed = myString[1]

	// %v is the utf-8 encoding, %T is the minimum datatype to hold it
	// this is because utf-8 has variable bit length
	fmt.Printf("%v, %T\n", indexed, indexed)

	for i, v := range myString {
		fmt.Println(i, v) // prints utf-encoding of character
	}

	// this is why runes may be better, which are just numbers that represent unicode characters
	var myRune = []rune("résumé")
	fmt.Printf("my rune = %v\n", myRune)

	// half to do it this way to make sure each char is a string, not rune
	// myString2 := "hi lol" makes each character a rune, which has to be converted
	myString2 := []string{"h", "i", " ", "l", "o", "l"}
	catString := ""
	for i := range myString2 {
		catString += myString2[i] // concatenation of strings
	}
	fmt.Println(catString)
}
