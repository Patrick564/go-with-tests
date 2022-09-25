package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Dependency Injection is pass a general type (for example as io.Writer) as
// argument, now you can use different methods who implement `io.Writer interface`
// as parameters for that method or function.
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreetHandler)))
}
