package main 
import  "fmt"

func array1() {
//declaring array of string data type
		var Colours = [5] string {"Blue", "White", "Brown", "Red",  "Yellow"}
		var myColours string = "My Best colours are"

		var temp = [3] string {"Tony", "Collins", "James"} 

		fmt.Println(myColours, Colours)

		//Accessing specified data in an array
		fmt.Println("I choose ",Colours[1])

		//Changing the element of temp array
		temp[1]="Joy"
		fmt.Println("The element of the second array was change from Collins to ",temp[1])


		//checking the length temp array
		//fmt.Println("The length of temp array is ",len(temp))
		fmt.Println("The length of temp array is ",cap(temp))

}

func main() {

	//declaring two arrays arr1 and arr2 of int data type

	//Declaring inferred array. It is very useful because you can any number of data
	var arr1 = [...] int {1,2,3,4,5,6}

	//This array type will print only the 3 secified datas
	arr2 := [3] int {5,6,7}

	fmt.Println("Printing array one datas",arr1)
	fmt.Println("Printing array two datas",arr2)

	//calling the functons of array1 and array2
	array1()
	
} 
