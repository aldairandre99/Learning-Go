package main

import (
	"fmt"
)

func main(){

	// Go Arrays

	// Arrays are used to store multiple values of the same type in a single variable, instead of declaring separate variables for each value.

	// In Go, there are two ways to declare an array 

	var arr1 = [4]int{1,2,3,4}
	arr2 := [3]int{5,6,7}

	fmt.Printf("\nArray1 : %v\nArray : %v\n\n",arr1,arr2)


}