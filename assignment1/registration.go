package assignment1

type Registration struct {
	Name string
	Surname  string
	Age        int
	Login      string
	Password   string
}

// type Database struct {
// 	logins []Registration
// }

func (r *Registration) Register(d *Database) {
	reg := Registration{r.Name, r.Surname, r.Age, r.Login, r.Password}
	d.Logins = append(d.Logins, reg)
}
