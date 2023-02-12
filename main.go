package main

// import "fmt"
import (
	"fmt"

	a "github.com/AminaAmangeldi15/Go/assignment1"
)

func main() {
	// var ar [] a.Registration
	d := a.Database{}
	// r := a.Registration{Name: "Amina", Surname: "Amangeldi", Age: 18, Login: "l", Password: "p"}
	// r1 := assignment1.Registration{Name: "Jin", Password: "Kim", Age: 24, Login: "l1", Password: "p1"}
	// r.Register(d)
	// response1 := d.Register(r)
	// fmt.Println(d)
	// fmt.Println(response1)
	// r1.Register(d)
	au := a.Authorization{Login: "l", Password: "p"}
	response := au.SignIn(d)
	fmt.Println(response)
}
