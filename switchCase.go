package main
import "fmt"

//Using switch case "int" and "string" data type

//Using "string" data type in switch case
func myColours()  {
	colour := "yellow"

	switch colour {

	case "blue":
	fmt.Println("colour is blue")

	case "red":
	fmt.Println("colour is red")

	case "white":
	fmt.Println("colour is white")

	case "yellow":
	fmt.Println("colour is yellow")

	default:
	fmt.Println("No colour match...")
	}


}


func myDays() {
//Using "int" data type in switch case
day := 5
	
switch day {
case 1:
fmt.Println("The Day is one")

case 2:
fmt.Println("The Day is two")

case 3:
fmt.Println("The Day is Three")

case 4:
fmt.Println("The Day is Four")

case 5:
fmt.Println("The Day is Five")

default:
fmt.Println("No day found...")

}


}

func MultiSwitchCase() {

	//using Multi switch case
	day := 5
   	switch day {

   	case 1,3,5:
    fmt.Println("Odd weekday")
   
	case 2,4:
     fmt.Println("Even weekday")
   	
	case 6,7:
    fmt.Println("Weekend")
  
	default:
    fmt.Println("Invalid day of day number")
  }


}

func main() {
	
		//calling the functions
		myColours()
		myDays()
		MultiSwitchCase()

}
