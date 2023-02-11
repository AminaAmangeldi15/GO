package registration

// import "github.com/aminaamangeldi15/go/assignment1/database"
// import "Go/assignment1"
// import "C:\Users\amina\KBTU\Go\assignment1\database.go"
// import (
// 	"./assignment1"
// )

type Registration struct {
	first_name string
	last_name  string
	age        int
	login      string
	password   string
}

// type Database struct {
// 	logins []Registration
// }

func (r *Registration) regis(d Database) {
	reg := Registration{r.first_name, r.last_name, r.age, r.login, r.password}
	d.logins = append(d.logins, reg)
}
