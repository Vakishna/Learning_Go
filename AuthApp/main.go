package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vakishna/go_crash_course/AuthApp/routes/callback"
	"github.com/vakishna/go_crash_course/AuthApp/routes/login"
	"github.com/vakishna/go_crash_course/AuthApp/routes/logout"
	"github.com/vakishna/go_crash_course/AuthApp/routes/user"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/callback", callback.CallBackHandler)
	r.HandleFunc("/login", login.LoginHandler)
	r.HandleFunc("/logout", logout.LogoutHandler)
	r.Handle("/user", negroni.New(
		negroni.HandlerFunc(middlewares.IsAuthenticated),
		negroni.Wrap(http.HandlerFunc(user.UserHandler)),
	))
	fmt.Println("Initialied Application :")
}
