package library

//User information
type User struct {
	ID 			int		`json:"id"`
	Username 	string	`json:"username"`
	Password 	string 	`json:"password"`
	Catagory    string  `json:"catagory"`
}

var (
	AllUsers []User
)
func init() {
	admin := User{
		Username: "admin",
		Password: "admin",
		Catagory: "admin",
	}
	user := User{
		Username: "user",
		Password: "user",
		Catagory: "user",
	}

	AllUsers = append(AllUsers, admin, user)
}