package playground

import "fmt"

func ArrayAsKeyValInMaps() {
	myArray := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	myArray1 := [10]int{100, 20, 3, -4, 5}
	myArray2 := [10]int{0, -2, -2}

	myMap := make(map[[10]int]bool)

	myMap[myArray] = true
	myMap[myArray1] = true
	myMap[myArray2] = true

	val, exist := myMap[myArray]

	if exist {
		fmt.Println(val)
	}
	// so this means that arrays are comparables

	if myArray != myArray1 {
		fmt.Println("they are comparable TOO")
	}
}
