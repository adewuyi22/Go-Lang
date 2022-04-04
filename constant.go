package main 
import  ("fmt")

func usingConstant() {

	const Height int = 100
	const Width int = 200
	
	var result int = Height * Width

	fmt.Print("The result = ", result)


}
func main () {

usingConstant()

}