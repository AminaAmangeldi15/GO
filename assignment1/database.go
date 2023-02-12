package assignment1

import "fmt"

// import "github.com/aminaamangeldi15/go/assignment1"
// import (
// 	"./assignment1"
// )
// import "C:\Users\amina\KBTU\Go\assignment1\registration.go"

type Database struct {
	Logins []Registration
}

func (d *Database) Register(r *Registration) *Registration {
	for i := 0; i < len(d.Logins); i++ {
		if d.Logins[i].Login == r.Login && d.Logins[i].Password == r.Password {
			// return fmt.Errorf("username %s %s already exists", r.Name, r.Surname)
			fmt.Println("User %s %s already exists", r.Name, r.Surname)
		}
	}
	reg := &Registration{r.Name, r.Surname, r.Age, r.Login, r.Password}
	d.Logins = append(d.Logins, Registration{r.Name, r.Surname, r.Age, r.Login, r.Password})
	fmt.Println("You registred!")
	return reg
}
