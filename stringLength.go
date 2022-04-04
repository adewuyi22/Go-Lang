package main 
import  ("fmt"
"reflect"

)

func length() {

	Fullname:= "Adewuyi Babajide James"
	fmt.Println("The number of length is: ", len(Fullname))
}


func main() {	

	length()

	var x string = "Welcome to Golang"
	fmt.Println(x)
	
	//Printing the Datatype
	fmt.Println("The Data type  is: ", reflect.TypeOf(x))
} 
