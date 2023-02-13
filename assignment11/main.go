package main

// import "fmt"
import (
	"fmt"
	"strings"
	// a "github.com/AminaAmangeldi15/Go/assignment1"
)

type Registration struct {
	Name     string
	Surname  string
	Age      int
	Login    string
	Password string
}

type Authorization struct {
	Login    string
	Password string
}

type Item struct {
	Name       string
	Price      float64
	Rating     float64
	RatingList []float64
}

type Database struct {
	Logins []Registration
	Items  []Item
}

func (d *Database) Register(r Registration) *Registration {
	for i := 0; i < len(d.Logins); i++ {
		if d.Logins[i].Login == r.Login && d.Logins[i].Password == r.Password {
			// return fmt.Errorf("username %s %s already exists", r.Name, r.Surname)
			fmt.Printf("User %s %s already exists!\n", r.Name, r.Surname)
			return nil
		}
	}
	reg := &Registration{r.Name, r.Surname, r.Age, r.Login, r.Password}
	d.Logins = append(d.Logins, Registration{r.Name, r.Surname, r.Age, r.Login, r.Password})
	fmt.Println("You registred!")
	return reg
}

func (a *Authorization) SignIn(d Database) string {
	for i := 0; i < len(d.Logins); i++ {
		if d.Logins[i].Login == a.Login && d.Logins[i].Password == a.Password {
			return "You entered system!"
		}
	}
	return "No authorized!!!"
}

func (d *Database) AddItem(name string, price float64) *Item {
	item := &Item{name, price, 0, nil}
	d.Items = append(d.Items, Item{name, price, 0, nil})
	return item
}

func (d *Database) SearchItem(name string) {
	for i := 0; i < len(d.Items); i++ {
		if strings.Contains(strings.ToUpper(d.Items[i].Name), strings.ToUpper(name)) {
			fmt.Println(d.Items[i])
			return
		} else {
			fmt.Printf("Did not find an item like %s", name)
		}
	}
}

func (d *Database) FilteringItems(price1, price2, rating1, rating2 float64) {
	for i := 0; i < len(d.Items); i++ {
		if d.Items[i].Price >= price1 && d.Items[i].Price <= price2 && d.Items[i].Rating >= rating1 && d.Items[i].Rating <= rating2 {
			fmt.Println("We found item that you searched!")
			fmt.Println(d.Items[i])
			return
		} else {
			fmt.Println("No such item with these price and rating!!!")
		}
	}
}

func Rate(d Database, a Authorization, itemName string, rating float64) {
	if a.SignIn(d) == "You entered system!" {
		fmt.Println("You entered system!")
		var sum float64
		for i := 0; i < len(d.Items); i++ {
			if d.Items[i].Name == itemName {
				d.Items[i].RatingList = append(d.Items[i].RatingList, rating)
				fmt.Println("You successfully rated an item!")
				for j := 0; j < len(d.Items[i].RatingList); j++ {
					sum += d.Items[i].RatingList[j]
				}
				d.Items[i].Rating = sum / float64(len(d.Items[i].RatingList))
				// fmt.Println(d.Items[i].Rating)
			} else {
				fmt.Println("No such item!!!")
			}
		}
	} else {
		fmt.Println("Wrong login or password!!!")
	}
}

func main() {
	// var ar [] a.Registration
	d := Database{}
	d.AddItem("Phone", 150000)
	r := Registration{Name: "Amina", Surname: "Amangeldi", Age: 18, Login: "l", Password: "p"}
	// r1 := Registration{Name: "Jin", Password: "Kim", Age: 24, Login: "l1", Password: "p1"}
	r1 := Registration{"Jin", "Kim", 24, "l1", "p1"}
	d.Register(r)
	d.Register(r1)
	fmt.Println(d)
	// r1.Register(d)
	au := Authorization{Login: "l", Password: "p"}
	response := au.SignIn(d)
	fmt.Println(response)
	au1 := Authorization{Login: "l1", Password: "p1"}
	response1 := au1.SignIn(d)
	fmt.Println(response1)
	// au2 := Authorization{Login: "l1", Password: "e"}
	// response2 := au2.SignIn(d)
	// fmt.Println(response2)
	// // d.Register(r)
	d.SearchItem("Asus")
	Rate(d, au, "Phone", 6)
	Rate(d, au1, "Phone", 7)
	fmt.Println(d.Items)
	// Rate(d, au, "Phone", 10)
	// d.FilteringItems(70000, 200000, 3, 7)
	// fmt.Println(d.Items)
}
