package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	//fmt.Printf("Hello, %s", name) this code writes the name to stdout, we need it to print to buffer
	fmt.Fprintf(writer, "Hello, %s", name)

}

//write is a great general purpose interface for "put this data somewhere"
//printf is very similar to Fprint, the difference is printf defaults to stdout wheras Fprintf
//takes a witer to send the string to.

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	//this is an http handler
	Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}

//in this example the output stream was the dependency, where the greet function could only greet to stdout
// we decided to inject dependecy, through parameters into the function. As we are passing interface to Fprintf, we can change dependecy to any other
// type until and unless it implements that interface.
