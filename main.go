package main

import (
    "fmt"
    "log"
    "net/http" //basic http server
)
// http ->module
func main() {
    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {  // /hello --> hello world 
        fmt.Fprintf(w, "Hello, World!")
    })
	http.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {   
        fmt.Fprintf(w, "DHARSHINI")
    })
	http.HandleFunc("/regnumber", func(w http.ResponseWriter, r *http.Request) {  
        fmt.Fprintf(w, "7376212IT129")
    })
	http.HandleFunc("/department", func(w http.ResponseWriter, r *http.Request) {  
        fmt.Fprintf(w, "IT")
    })
	http.HandleFunc("/color", func(w http.ResponseWriter, r *http.Request) {   
        fmt.Fprintf(w, "Black")
    })
    fmt.Printf("Server running (port=8080), route: http://localhost:8080/helloworld\n")
    if err := http.ListenAndServe(":8080", nil); err != nil { // port number 80:8080 
        log.Fatal(err) //throw error 
    }
}