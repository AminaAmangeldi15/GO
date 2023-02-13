package main

// import "fmt"
import (
	"fmt"

	a "github.com/AminaAmangeldi15/Go/assignment1"
)

func main() {
	// var ar [] a.Registration
	d := a.Database{}
	// i:= a.Item{"Phone", 150000, 0}
	
	// d.AddItem(a.Item{"Phone", 150000, 0})
	r := a.Registration{Name: "Amina", Surname: "Amangeldi", Age: 18, Login: "l", Password: "p"}
	// r1 := Registration{Name: "Jin", Password: "Kim", Age: 24, Login: "l1", Password: "p1"}
	r1 := a.Registration{Name: "Jin", Surname: "Kim", Age: 24, Login: "l1", Password: "p1"}
	// d.Register(r)
	// d.Register(r1)
	r.Register(d)
	r1.Register(d)
	fmt.Println(d)

	au := a.Authorization{Login: "l", Password: "p"}
	// response := au.SignIn(d)
	// fmt.Println(response)
	au.SignIn(d)
	// au1 := Authorization{Login: "l1", Password: "p"}
	// response1 := au1.SignIn(d)
	// fmt.Println(response1)
	// d.Register(r)
	// d.SearchItem("Phone")
	// Rate(d, au, "Phone", 6)
	// fmt.Println(d.Items)
	// d.FilteringItems(70000, 200000, 3, 7)
}