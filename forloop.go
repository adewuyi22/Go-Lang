package main
import "fmt"

var num int = 0
var me string = "me"

func counting() {
	for num<=5 {
		fmt.Println("Looping", me, num)
		num++
	}
}

//using continue statement
func cont() {
	for i :=0; i<5; i++ {
		if i==3 {
			continue

			//fmt.Println("Reached three")
		}
		fmt.Println(i)
	}


}

//using nested loop
func nest() {
	adj := [2]string{"big", "tasty"}
	fruits := [3]string{"apple", "orange", "banana"}
	for i:=0; i < len(adj); i++ {
	  for j:=0; j < len(fruits); j++ {
		fmt.Println(adj[i],fruits[j])
	  }
	}

}

func main() {
	//calling the function
	counting()
	cont()
	nest() 
	
}
	