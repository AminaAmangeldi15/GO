package assignment1

type Authorization struct {
	Login    string
	Password string
}

// func (a *Authorization) SignIn(d Database) string {
// 	for i := 0; i < len(d.Logins); i++{
// 		if d.Logins[i].Login == a.Login && d.Logins[i].Password == a.Password{
// 			return "You entered system!"
// 		}
// 	}
// 	return "No authorized!!!"
// }

func (a *Authorization) SignIn(d Database) string {
	for i := 0; i < len(d.Logins); i++ {
		if d.Logins[i].Login == a.Login && d.Logins[i].Password == a.Password {
			return "You entered system!"
		}
	}
	return "No authorized!!!"
}
