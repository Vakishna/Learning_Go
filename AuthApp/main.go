package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"

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
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
	http.Handle("/", r)
	log.Print("Server listen")
	log.Fatal(http.ListenAndServe(":3000", nil))

	fmt.Println("Initialied Application :")
}
