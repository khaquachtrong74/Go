package main

import (
	"flag"
	"fmt"
)
func main(){
	fmt.Println("Hello World!")
	var numbers = flag.Int("n", 5124, "help message for flag n")
	var B  = flag.Bool("b", true, "A bool")

	var Str  = flag.String("s","A foo","string here!")
	flag.Parse()
	fmt.Println(*numbers)
	fmt.Println(*B)
	fmt.Println(*Str)
}
