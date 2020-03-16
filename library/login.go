package library

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// CheckAuth takes http.Request as parameter and checks requests's authorization
// header. For invalid username/password, it returns error & a flag
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in check")
	username, password, authErr := r.BasicAuth()
	fmt.Println(username, password)
	if authErr == false {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Couldn't parse login information correctly!!"))
		return
	}
	//var catagory string
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user.Catagory, "-------------")
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, user := range AllUsers {
		if user.Username == username && user.Password == password {
			fmt.Println("user found")
			token, err := GenerateToken(username, user.Catagory)
			if err != nil {
				w.Write([]byte("Error generating token"))
			} else {
				w.Write([]byte(token))
			}
		} else if user.Username == username && user.Password != password {
			w.Write([]byte("invalid password"))
		}
	}
}
