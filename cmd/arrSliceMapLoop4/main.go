package main

import "fmt"

func main() {
	arrays()
	slices()
	maps()
	loops()
}

func arrays() {
	var intArr [3]int32      // 3-element int32 array
	fmt.Println(intArr[0])   // defaults to 0
	fmt.Println(intArr[1:3]) // like python
	intArr[1] = 2
	fmt.Println(intArr)

	// arrays are stored contiguously in memory
	// for a int32 array, each element is 4 memory spots away, since 32-bit = 4 bytes
	fmt.Println(&intArr) // memory address is & like C
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	var intArr2 [3]int32 = [3]int32{1, 2, 3} // declare and initialize at the same time
	fmt.Println(intArr2)

	intArr4 := [3]int32{1, 2, 3}
	fmt.Println(intArr4)

	intArr5 := [...]int32{1, 2, 3} // compiler infers size as fixed size 3
	fmt.Println(intArr5)
}

func slices() {
	// wrapper around arrays, giving ability to make non-fixed sizes
	var intSlice []int32 = []int32{1, 2, 3} // array of size 3 is made
	fmt.Println(intSlice)
	fmt.Printf("slice len before append is %v and capacity is %v\n", len(intSlice), cap(intSlice))

	// attempt to append to slice. if not enough space, create a new array with enough space
	intSlice = append(intSlice, 4) // elements passed "4" are not declared yet and can't be read
	fmt.Println(intSlice)
	fmt.Printf("slice len after append is %v and capacity is %v\n", len(intSlice), cap(intSlice))
	// you can omit the type after variable name for slices, maps, arrays like so
	var intSlice2 = []int32{5, 6}
	intSlice = append(intSlice, intSlice2...) // append multiple elements
	fmt.Println(intSlice)

	// new int32 slice of capacity 8 and first 3 elemtns declared to 0
	// best to initialize with a preset capacity because reallocation of slice takes longer
	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)
}

func maps() {
	var myMap map[string]uint8 = make(map[string]uint8) // {string: unsigned int}
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"adam": 23, "Sarah": 45} // declare and initialize
	myMap2["bob"] = 32
	fmt.Println(myMap2["bob"])

	// beware! maps return a value whether or not the key is in the map
	// always check the existence value of the return
	age, exist := myMap2["john"]
	if exist {
		fmt.Printf("they are %v years old\n", age)
	} else {
		fmt.Println("bruh")
	}

	delete(myMap2, "bob")
	fmt.Println(myMap2)
}

func loops() {
	var intArr = [5]int32{4, 3, 9, 6, 4}
	for i, v := range intArr {
		fmt.Printf("%v index: %v\n", i, v)
	}

	var myMap = map[string]uint8{"adam": 32, "bailey": 45}
	for name, age := range myMap {
		fmt.Printf("%v is %v years old \n", name, age)
	}

	// go does not explicitly have a while loop. it has this instead
	i := 1
	for {
		if !(i < 11) {
			break
		}
		fmt.Println(i)
		i++ // go does not have ++i, only i++
	}

	// does same thing as above
	for j := 1; j < 11; j++ {
		fmt.Println(j)
	}
}
