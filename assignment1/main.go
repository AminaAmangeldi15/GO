package main

// import "fmt"
import (
	"fmt"

	// a "./assignment1"
)

func main(){
	d := Database{}
	r := Registration{Name: "Amina", Surname: "Amangeldi", Age: 18, Login: "l", Password: "p"}
	r1 := Registration{Name: "Jin", Password: "Kim", Age: 24, Login: "l1", Password: "p1"}
	r.Register(d)
	r1.Register(d)
	au := Authorization{Login: "l", Password: "p"}
	response := au.SignIn(d)
	fmt.Println(response)
}
