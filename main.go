package main

import (
	"fmt"
)

func main(){

	// Syntax

	// const CONSTNAME type = value

	// Note: The value of a constant must be assigned when you declare it. 

	// Multiple Constants Declaration

	// Multiple constants can be grouped together into a block for readability:

	const YEAR_BIRTH = 1999
	const (
		MONTH = 12 
		DAY = 30
	)

	fmt.Printf("\nBirthday Year: %v\nMonth: %v\nDay: %v\n\n",YEAR_BIRTH,MONTH,DAY)

}