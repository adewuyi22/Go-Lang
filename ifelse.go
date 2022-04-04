package main
import "fmt"

func BioData() {
var FirstName string = "Adex"
var LastName string = "James"

if (FirstName == "Adex") {
	fmt.Println("First Name is", FirstName)
} else {
	fmt.Println("No Adex found")
}

if (LastName == "James") {
	fmt.Println("Last Name is",  LastName)
}

}


func main(){
	BioData()
}	