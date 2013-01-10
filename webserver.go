package main
import "fmt"
import "net/http"
import "io"

//hämta mätvärde från arduinon
//visa upp mätvärde


func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main (){
	fmt.Println("hello")
	err := http.ListenAndServe(":12345", http.HandlerFunc(HelloServer))
	if err != nil {
		panic(err)
	}
	
	
}
