package main

// import "fmt"
import (
	"fmt"

	a "github.com/AminaAmangeldi15/Go/assignment1/a1"
)

func main() {
	// var ar [] a.Registration
	d := a.Database{}
	d.AddItem("Phone", 150000)
	r := a.Registration{Name: "Amina", Surname: "Amangeldi", Age: 18, Login: "l", Password: "p"}
	// r1 := Registration{Name: "Jin", Password: "Kim", Age: 24, Login: "l1", Password: "p1"}
	r1 := a.Registration{"Jin", "Kim", 24, "l1", "p1"}
	d.Register(r)
	d.Register(r1)
	fmt.Println(d)
	// r1.Register(d)
	au := a.Authorization{Login: "l", Password: "p"}
	response := au.SignIn(d)
	fmt.Println(response)
	au1 := a.Authorization{Login: "l1", Password: "p1"}
	response1 := au1.SignIn(d)
	fmt.Println(response1)
	// au2 := Authorization{Login: "l1", Password: "e"}
	// response2 := au2.SignIn(d)
	// fmt.Println(response2)
	// // d.Register(r)
	d.SearchItem("Asus")
	a.Rate(d, au, "Phone", 6)
	a.Rate(d, au1, "Phone", 7)
	fmt.Println(d.Items)
	// Rate(d, au, "Phone", 10)
	// d.FilteringItems(70000, 200000, 3, 7)
	// fmt.Println(d.Items)
}
