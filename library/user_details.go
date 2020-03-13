package library

import (
	"errors"
	"fmt"
	"net/http"
)

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


// CheckAuth takes http.Request as parameter and checks requests's authorization
// header. For invalid username/password, it returns error
func checkAuth(r *http.Request) error {
	fmt.Println("in check")
	username, password, ok := r.BasicAuth()
	if !ok {
		return errors.New("unauthorized")
	}

	for _, user := range AllUsers {
		if user.Username == username && user.Password == password {
			fmt.Println("user found")
			if user.Catagory != "admin"{
				return errors.New("You are not an admin")
			}
			return nil
		} else if user.Username == username && user.Password != password {
			return errors.New("invalid password")
		}
	}
	return errors.New("user not found")
}

func AuthMiddleware(next http.Handler) http.Handler {
	fmt.Println("in auth")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := checkAuth(r); err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}