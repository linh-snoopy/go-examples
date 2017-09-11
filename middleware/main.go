package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"net/http"
	"sync"
	"crypto/subtle"
)

type helloWorldHandler struct {
}

var sessionStore map[string]Client
var storageMutex sync.RWMutex

type Client struct {
	loggedIn bool
}

type authenticationMiddleware struct {
	wrappedHandler http.Handler
}

const loginPage = `
<html>
	<head>
		<title>Login</title>
	</head>
	<body>
		<form action="/login" method="post"> 
			<input type="password" name="password" /> 
			<input type="submit" value="login" /> 
		</form> 
	</body> 
</html>`

func main() {
	sessionStore = make(map[string]Client)

	http.Handle("/hello", helloWorldHandler{})
	http.Handle("/secureHello", authenticate(helloWorldHandler{}))
	http.HandleFunc("/login", handleLogin)

	http.ListenAndServe(":3000", nil)
}

func (h helloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!!!")
}

func (h authenticationMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		if err != http.ErrNoCookie {
			fmt.Fprint(w, err)
		} else {
			err = nil
		}
	}
	var present bool
	var client Client
	if cookie != nil {
		storageMutex.RLock()
		client, present = sessionStore[cookie.Value]
		storageMutex.RUnlock()
	} else {
		present = false
	}

	if !present {
		cookie = &http.Cookie{
			Name:  "session",
			Value: uuid.NewV4().String(),
		}
		client = Client{false}
		storageMutex.RLock()
		sessionStore[cookie.Value] = client
		storageMutex.RUnlock()
	}

	http.SetCookie(w, cookie)
	// if haven't logged in yet, route to login page
	if !client.loggedIn {
		fmt.Fprint(w, loginPage)
		return
	} else {
		h.wrappedHandler.ServeHTTP(w, r)
	}
}

func authenticate(h http.Handler) authenticationMiddleware {
	return authenticationMiddleware{h}
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		if err != http.ErrNoCookie {
			fmt.Fprint(w, err)
		} else {
			err = nil
		}
	}
	var present bool
	var client Client
	if cookie != nil {
		storageMutex.RLock()
		client, present = sessionStore[cookie.Value]
		storageMutex.RUnlock()
	} else {
		present = false
	}

	if present == false {
		cookie = &http.Cookie{
			Name:  "session",
			Value: uuid.NewV4().String(),
		}
		client = Client{false}
		storageMutex.Lock()
		sessionStore[cookie.Value] = client
		storageMutex.Unlock()
	}
	http.SetCookie(w, cookie)

	err = r.ParseForm()
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	if subtle.ConstantTimeCompare([]byte(r.FormValue("password")), []byte("password123")) == 1 {
		client.loggedIn = true
		fmt.Fprintln(w, "Thank you for logging in.")
		storageMutex.Lock()
		sessionStore[cookie.Value] = client
		storageMutex.Unlock()
	} else {
		fmt.Fprintln(w, "Wrong password")
	}
}
