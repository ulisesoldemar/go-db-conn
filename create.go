package main

import "net/http"

func create(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "create", nil)
}

func insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		email := r.FormValue("email")
		connection := connectionDB()
		insertReg, err := connection.Prepare("INSERT INTO employee(name, email) VALUES (?, ?)")
		if err != nil {
			panic(err.Error())
		}
		insertReg.Exec(name, email)

		http.Redirect(w, r, "/", 301)
	}
}
