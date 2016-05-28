package router

import (
	"fmt"
	"net/http"
)

type LoginRouter struct {
	SrvMux *http.ServeMux
}

// Register is register Router
func (lr LoginRouter) Register() {
	lr.SrvMux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "You are in [%v].", "Register")
	})
}

// Login is login Router
func (lr LoginRouter) Login() {
	lr.SrvMux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "You are in [%v].", "Login")
	})
}

// Logout is logout Router
func (lr LoginRouter) Logout() {
	lr.SrvMux.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "You are in [%v].", "Logout")
	})
}
