package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm() // default not parse
		fmt.Println(r.Form)
		fmt.Println("path:", r.URL.Path)
		fmt.Println("Scheme:", r.URL.Scheme)
		fmt.Println(r.Form.Get("url_long"))

		for k, v := range r.Form {
			fmt.Println("key:", k)
			fmt.Println("val:", strings.Join(v, ""))
		}

		fmt.Fprintf(w, "Hello Weixin.")
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request method: ", r.Method)

		if r.Method == "GET" {
			t, _ := template.ParseFiles("login.gtpl")
			t.Execute(w, nil)
		} else {
			r.ParseForm()
			uname := template.HTMLEscapeString(r.Form.Get("username"))
			if m, _ := regexp.MatchString("^[a-zA-Z0-9]*$", uname); !m {
				fmt.Println("[" + uname + "] must be number mixed ascii")
				fmt.Fprintf(w, "uname must be number or ascii or mixed two of them")
				return
			}

			fmt.Println("username:", uname)
			fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))

			template.HTMLEscape(w, []byte(r.Form.Get("username")))
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln("Start server at port 8080 failed!")
	}
}
