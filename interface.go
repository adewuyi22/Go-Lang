package main 
import  "fmt"

/*

Go has different approaches to implement the concepts of object-orientation. 
Go does not have classes and inheritance. Go fulfill these requirements 
through its powerful interface. Interfaces provide behavior to an object: 
if something can do this, then it can be used here.
An interface defines a set of abstract methods and does not contain any variable.

*/

//Declaring an interface (object)
type vehicle interface {
	accelerate()
}

//Declaring a function
func foo (v vehicle) {
	fmt.Println(v)
}


type car struct {
	model string
	color string
}

func (c car) accelerate() {
	fmt.Println("Accelrating...")
}
	
type toyota struct {
	model string
	color string
	speed int
}

func (t toyota) accelerate() {
	fmt.Println("I am toyota, I accelerate fast?")
}


func main() {	
c1:=car {"Suzuki", "blue"}
t1:=toyota{"Toyota", "Red", 100}
c1.accelerate()
t1.accelerate()
foo(c1)
foo(t1)

} 
