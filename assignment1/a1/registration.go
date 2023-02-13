package a1

// import "fmt"

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

// func (r *Registration) Register(d *Database) error{
// 	for i := 0; i < len(d.Logins); i++{
// 		if d.Logins[i].Login == r.Login && d.Logins[i].Password == r.Password{
// 			return fmt.Errorf("username %s %s already exists", r.Name, r.Surname) 
// 		} 
// 	}
// 	reg := Registration{r.Name, r.Surname, r.Age, r.Login, r.Password}
// 	d.Logins = append(d.Logins, reg)
// 	return nil



// }
