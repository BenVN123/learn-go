package main

import "fmt"
import "errors"

/*
func funcName(param1 datatype1, param2 datatype2) (output1, output2) {
    return param1, param2
}
*/

func main() {
	var str string = "hi lol"
	printMe(str)

	var integer int
	var err error
	integer, err = intDivision(123, 0)
	if err != nil { // error handling is done like this
		fmt.Println(err.Error())
	} else {
		fmt.Println(integer)
	}

	switch { // same thing as the if above
	case err != nil:
		fmt.Println(err.Error())
	default:
		fmt.Println(integer)
	}

	switch integer { // more switch
	case 1:
		printMe("1 lol")
	case 2, 3:
		printMe("more lol")
	default:
		printMe("nothing lol")
	}

}

func printMe(value string) {
	fmt.Println(value)
}

func intDivision(num1 int, num2 int) (int, error) {
	var err error
	if num2 == 0 {
		err = errors.New("cant divide by zero")
		return 0, err
	}
	return num1 / num2, err
}
