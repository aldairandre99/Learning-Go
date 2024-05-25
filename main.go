package main

import (
	"fmt"
)

func main(){

	// Formatting Verbs for Printf()

	// Go offers several formatting verbs that can be used with the Printf() function.


	/* 
		
		%v 	Prints the value in the default format
		%#v 	Prints the value in Go-syntax format
		%T 	Prints the type of the value
		%% 	Prints the % sign
		
		
	*/

	var i = 15.5
  var txt = "Hello World!"

  fmt.Printf("%v\n", i)
  fmt.Printf("%#v\n", i)
  fmt.Printf("%v%%\n", i)
  fmt.Printf("%T\n", i)

  fmt.Printf("%v\n", txt)
  fmt.Printf("%#v\n", txt)
  fmt.Printf("%T\n", txt) 


}