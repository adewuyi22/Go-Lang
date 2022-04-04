package main 
import  (
	"fmt"
	"regexp"
)

func findString() {
	//This will search and display any/list of text that has ".org"
	re := regexp.MustCompile(".org")
	
	fmt.Println(re.FindString("google.com"))
	fmt.Println(re.FindString("abc.org"))
	fmt.Println(re.FindString("fb.org"))
}


func findIndex() {
	// returns a string having the text of the left most match.

	fmt.Println("\n")
	re := regexp.MustCompile(".org")
	
	fmt.Println(re.FindStringIndex("google.com"))
	fmt.Println(re.FindStringIndex("abc.org"))
	fmt.Println(re.FindStringIndex("fb.org"))
}

func submatch() {
	//This will search for a matching text starting with "f" and ing
	fmt.Println("\n")
	re := regexp.MustCompile("f([a-z]+)ing")
	
	fmt.Println(re.FindStringSubmatch("flying1"))
	fmt.Println(re.FindStringSubmatch("abcfloatingingxyz"))
}

func main() {	
	findString()
	findIndex()
	submatch()
} 
