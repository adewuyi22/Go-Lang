package main 
import  ("fmt")

//Declaring a global variable
var usingFunctions string = "Trying to use functions in Go Lang"

//first function
func show1() {
	var firstFunction = "This is my 1st function"
	fmt.Println(firstFunction)
}

//second function
func show2() {
	var secondFunction = "This is my 2nd function"
	fmt.Println(secondFunction) 
}

func main() {

	fmt.Println(usingFunctions)

	//calling the two functions
	show1()
	show2()
} 
