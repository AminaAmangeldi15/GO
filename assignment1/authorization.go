package assignment1

// import "github.com/aminaamangeldi15/go/assignment1/database"
// import "Go/assignment1"2

type Authorization struct {
	Login    string
	Password string
}

func (a *Authorization) SignIn(d *Database) string {
	for i := 0; i < len(d.logins); i++{
		if d.logins[i].Login == a.Login && d.logins[i].Password == a.Password{
			return "You entered system!"
		} 
	}
	return "No authorized!!!"
}
