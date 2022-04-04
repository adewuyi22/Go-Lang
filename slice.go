package main 
import  "fmt"

/*

	Slices are similar to arrays, but are more powerful and flexible.

	Like arrays, slices are also used to store multiple values of the same type in a single variable.

	However, unlike arrays, the length of a slice can grow and shrink as you see fit.

	In Go, there are several ways to create a slice:

	Using the []datatype{values} format
	Create a slice from an array
	Using the make() function

*/

	func data1()  {
	//This will print 2, 4, 6 from the slice element 
	var odd = [6] int {2,4,6,8,10,12}
	var s [] int = odd[0:3]
	var s1 [] int = odd[0:4]
	var s2 [] int = odd[0:5]

	fmt.Println("Printing Data 1 details: \n", s)
	fmt.Println(s1)
	fmt.Println(s2)
	}

	func data2() {
	var names = [4] string {"James", "Tina", "Jack", "Ben",} 
	var name2 [] string = names [0:3]  

	fmt.Println("\nPrinting Data 2 details: \n", names)
	fmt.Println(name2)
	
	
	//var slice1 = names[0:2]
	//var slice2 = names [1:3]

	//slice2[0] = "ZZZ"
	//fmt.Println(slice1, slice2)
	//fmt.Println(names)

 }

func main() {	
	data1()	
	data2()	
	
} 
