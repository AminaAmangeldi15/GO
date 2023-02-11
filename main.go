package main

// import "fmt"
import (
	a "./assignment1"
)

func main(){
	d := a.Database{}
	first_name := "Amina"
	last_name := "Amangeldi"
	age := 18
	l := "amina"
	p := "pass"
	reg := a.Registration{first_name, last_name, age, l, p}
	reg.register(d)
}
