package main

import "fmt"

import "unicode/utf8"

func main() {
	var intNum int = 32767 // could be int32 or int64 depending on machine
	var uIntNum uint = 123 // unsigned int
	intNum = intNum + 1
	intNum += 1
	// cant mix types at all, must cast (following is error): fmt.Println(intNum + uIntNum)
	fmt.Println(intNum + int(uIntNum))

	var floatNum float32 = 123.9 // 32-bit signed float
	fmt.Println(floatNum + float32(123))

	var intNum1 int32 = 3
	var intNum2 int32 = 2
	fmt.Println(intNum1 / intNum2)                   // truncates division for ints, like C
	fmt.Println(float32(intNum1) / float32(intNum2)) // cast first to get correct number

	var myString string = "Hello \nWorld" // double quotes are strings
	var myString2 string = `Hello 
World` // backticks are multi line strings, like python
	fmt.Println(myString)
	fmt.Println(myString2)

	fmt.Println(len("test")) // len() returns number of bytes, watch out when using unicode!
	fmt.Println(len("∑"))    // unicode will show extra bytes. len only works with ascii

	fmt.Println(utf8.RuneCountInString("∑")) // for unicode, use Rune functions
	var myRune rune = '∑'
	fmt.Println(myRune) // cant print rune since rune stored as an integer code, convert to string

	var myBoolean bool = false // bool :)
	fmt.Println(myBoolean)

	var defaultInt int       // default value for all ints, floats, runes is 0
	var defaultBool bool     // default bool is false
	var defaultString string // default string is empty string
	fmt.Println(defaultInt)
	fmt.Println(defaultBool)
	fmt.Println(defaultString)

	// implicit datatypes
	var myVar = "text" // compiler predicts that it is string
	myVar2 := "text"   // same as previous
	fmt.Println(myVar, myVar2)

	var1, var2 := 1, 2 // assign multiple variables at a time
	fmt.Println(var1, var2)

	const myConst = "const string"
	fmt.Println(myConst)
}
