package library

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"strings"
	"time"
)


//Get JWT token for authorization
func GenerateToken(username string, cata string) (string, error){
	fmt.Println("generating token")

	mySigningKey := []byte("secretkey")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	fmt.Println(cata)
	claims["type"] = cata
	claims["user"] = username
	claims["exp"] = time.Now().Add(time.Minute * 300).Unix()
	token.Claims = claims
	fmt.Println(claims)
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "",err
	}
	fmt.Println("done generating token")
	return tokenString, nil
}

//func AuthMiddleware(next http.Handler) http.Handler {
//	fmt.Println("in auth")
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		if err, ok := checkAuth(r); !ok {
//			http.Error(w, err.Error(), http.StatusUnauthorized)
//		} else {
//			next.ServeHTTP(w, r)
//		}
//	})
//}


//to validate jwt tokens
func IsAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET"{
			log.Println(r.Method)
			next.ServeHTTP(w, r)
		}else{

			fmt.Println("Middleware")
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Need jwt-token for authorization! http://localhost:<port>/token\n", http.StatusUnauthorized)
				return
			}

			token, err:= jwt.Parse(strings.Split(authHeader, " ")[1], func(token *jwt.Token) (interface{}, error) {
				return []byte("secretkey"), nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				fmt.Println("Here")
				fmt.Println(claims)
				fmt.Println(claims["type"].(string))
				if claims["type"].(string) != "admin" {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("you are not admin"))
					return
				}
				next.ServeHTTP(w, r)
			} else if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors & jwt.ValidationErrorMalformed != 0 {
					w.Write([]byte(fmt.Sprintf("That's not even a token\n")))
				} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
					// Token is either expired or not active yet
					w.Write([]byte(fmt.Sprintf("The token is Expired! Please issue a new token!")))
				} else {
					w.Write([]byte(fmt.Sprintf("Couldn't handle this token: ", err)))
				}
			} else {
				w.Write([]byte(fmt.Sprintf("Couldn't handle this token: ", err)))
			}
		}

	})
}