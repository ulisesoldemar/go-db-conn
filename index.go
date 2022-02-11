package main

import "net/http"

type Employee struct {
	Id    int
	Name  string
	Email string
}

// Esta función recepciona los datos
func index(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hi!")
	connection := connectionDB()
	regs, err := connection.Query("SELECT * FROM employee")
	if err != nil {
		panic(err.Error())
	}
	employee := Employee{}
	employees := []Employee{}

	for regs.Next() {
		var id int
		var name string
		var email string
		err := regs.Scan(&id, &name, &email)

		if err != nil {
			panic(err.Error())
		}
		employee.Id = id
		employee.Name = name
		employee.Email = email

		employees = append(employees, employee)
	}
	templates.ExecuteTemplate(w, "index", employees)
}
