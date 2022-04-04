package main
import 
"net/http"
"fmt"

func main() {
	http.HandleFunc("/", someFunc)
	http.ListenAndserve(": 8080", nil)
}

func someFunc(w http.ResponseWriter, req *http.Request) {
	w.Write([] byte ("hello Universe"))
}
