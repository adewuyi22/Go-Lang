package main

import "fmt"

//declaring variable name 'welcome'
var welcome string = "\nWelcome to Golang! \nThe Flexible and fun to code Language"
 

var j string = "\ajamezo"
var o string = "_official"

var c = j + o

var num int = 2 + 2*5

var num1 int = 5
var num2 int = 10

var a, b, d int = 100, 101, 102

var me bool = true

func main() {

	if me == true {
		fmt.Println("Boolean  is true")
	} else {
		fmt.Println("Boolean is false")
	}

	if num1 >= 5 {
		fmt.Println("Num1 is = 5")
	} else {
		fmt.Println("Not = 5")
	}

	//printing the variable type with %T while the value type is %V
	fmt.Printf("The variable type of num1 is %T and the value is %v\n", num, num)

	fmt.Println(a, b, d)
	fmt.Println(me)
	fmt.Println(num)
	fmt.Println(c)
	fmt.Println(welcome)

	
}
