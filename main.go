package main

// import "fmt"
import (
	"fmt"

	"github.com/AminaAmangeldi15/Go/assignment1"
)

func main() {
	d := assignment1.Database{}
	r := assignment1.Registration{Name: "Amina", Surname: "Amangeldi", Age: 18, Login: "l", Password: "p"}
	// r1 := assignment1.Registration{Name: "Jin", Password: "Kim", Age: 24, Login: "l1", Password: "p1"}
	r.Register(d)
	// r1.Register(d)
	au := assignment1.Authorization{Login: "l", Password: "p"}
	response := au.SignIn(d)
	fmt.Println(response)
}
