package main

import (
	"fmt"
)

func main(){

	//Syntax
	// var variablename type = value //type is string 

	// Or you can use 

	// var variablename = value /type is inferred
	// variablename := value /type is inferred only can be used insede functions

	var age = 24
	var name string = "Jonh"
	lastname := "Fake"

	fmt.Printf("\nHello...\nMy name is %v %v\nAnd i'm a %v year old based in Angola\n",name,lastname,age)
	println(" ")
}