package main

import (
	"fmt"
	"io"
	"net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request){
	fmt.Printf("got / request\n")
	io.WriteString(w, "It is a website\n")
}

func getGreeting(w http.ResponseWriter, r *http.Request){
	fmt.Printf("got /greeting request\n")
	io.WriteString(w, "Hello\n")
}

func main(){
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/greeting", getGreeting)
	fmt.Println("Server is listening...")
	err := http.ListenAndServe("8181", nil)
	if err != nil{
		return
	}
}