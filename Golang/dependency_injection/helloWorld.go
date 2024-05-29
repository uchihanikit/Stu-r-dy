package dependency_injection

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writer io.Writer, name string) {

	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(writer http.ResponseWriter, request *http.Request) {
	Greet(writer, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
