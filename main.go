package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexPage)
	http.HandleFunc("/contact", contactPage)
	http.ListenAndServe(":8080", nil)
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to Abhiram's site!</h1>")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "Email us at <a href=\"mailto:mail@odyn.com\">mail@odyn.com</a>.")
}
