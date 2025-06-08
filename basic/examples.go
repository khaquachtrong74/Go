package main

import (
	"fmt"
)

func main(){

	// single creation
	var singleVal int
	singleVal = 10
	fmt.Println("Single Value", singleVal)
	
	// compound creation
	var a, b = 1, "Coders"
	fmt.Println("compound creation\na: ", a,"\nb:", b)

	// block creation
	{
		fmt.Println("Start block")
		var blockVariable int
		blockVariable = 5
		fmt.Println(blockVariable)
		fmt.Println("End block")
	}
	// blockVariable is not accessible here
	// Error:  fmt.Println(blockVariable)
}
