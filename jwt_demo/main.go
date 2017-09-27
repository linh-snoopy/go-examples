package main

//Now start it up and visit /settoken to simulate loging in
//Then you’ll be redirected to /profile which is only accessible with your token
//Lastly simulate logging out by going to /logout

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"github.com/dgrijalva/jwt-go"
	"log"
	"strings"
)

type Key int

const MyKey Key = 0

type Claims struct {
    Username string `json:"username"`
    // recommended having
    jwt.StandardClaims
}

// create a JWT and put in the clients cookie
func setToken(res http.ResponseWriter, req *http.Request) {
	log.Println("setToken")
    // Expires the token and cookie in 1 hour
    expireToken := time.Now().Add(time.Second * 20).Unix()
    expireCookie := time.Now().Add(time.Second * 20)

    // We'll manually assign the claims but in production you'd insert values from a database 
    claims := Claims {
        "myusername",
        jwt.StandardClaims {
            ExpiresAt: expireToken,
            Issuer:    "localhost:9000",
        },
    }

    // Create the token using your claims
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Signs the token with a secret.    
    signedToken, _ := token.SignedString([]byte("secret"))
	log.Println(signedToken)

    // Place the token in the client's cookie 
    cookie := http.Cookie{Name: "Auth", Value: signedToken, Expires: expireCookie, HttpOnly: true}
    http.SetCookie(res, &cookie)
	log.Println("settoken cookie:", cookie)
    
    // Redirect the user to his profile
    http.Redirect(res, req, "/profile", 307)
}

// middleware to protect private pages
func validate(page http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		cookie, err := req.Cookie("Auth")
		if err != nil {
			log.Println(err)
			http.NotFound(res, req)
			return
		}

		token, err := jwt.ParseWithClaims(cookie.Value, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			log.Println("in function:", token)
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method")
			}
			return []byte("secret"), nil
		})
		log.Println("validate token:", token)
		if err != nil {
			log.Println(err)
			http.NotFound(res, req)
			return
		}

		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			ctx := context.WithValue(req.Context(), MyKey, *claims)
			page(res, req.WithContext(ctx))
		} else {
			http.NotFound(res, req)
			return
		}
	})
}

// only viewable if the client has a valid token
func protectedProfile(res http.ResponseWriter, req *http.Request){
	log.Println("protectedProfile")
    claims, ok := req.Context().Value(MyKey).(Claims)
    if !ok {
        http.NotFound(res, req)
        return
    }

    fmt.Fprintf(res, "Hello %s", claims.Username)
}

// deletes the cookie
func logout(res http.ResponseWriter, req *http.Request){
	log.Println("logout")
    deleteCookie := http.Cookie{Name: "Auth", Value: "none", Expires: time.Now()}
    http.SetCookie(res, &deleteCookie)
    return
}

func protectedProfile2(res http.ResponseWriter, req *http.Request) {
	autoken := req.Header.Get("Authorization")
	t := strings.Split(autoken, " ")
	token, err := jwt.ParseWithClaims(t[1], &Claims{}, func(token *jwt.Token) (interface{}, error) {
			log.Println("in function:", token)
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method")
			}
			return []byte("secret"), nil
		})
		log.Println("Authorization token:", token)
		if err != nil {
			log.Println(err)
			http.NotFound(res, req)
			return
		}
	res.Write([]byte("qqqqqqqqqqqqqqqqqqqqqq"))
}

func main(){
    http.HandleFunc("/settoken", setToken)
    http.HandleFunc("/profile", validate(protectedProfile))    
    http.HandleFunc("/logout", validate(logout))
	http.HandleFunc("/test", protectedProfile2)
	fmt.Println("Server is starting")
	log.Fatal(http.ListenAndServe(":9000", nil))
}