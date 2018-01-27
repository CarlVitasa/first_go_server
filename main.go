package main

import (	
	"fmt"
	"net/http"
	"text/template"
	// "log"
)

var defaultPerson = new(Person)

func main() {
	fmt.Println("Hello, World!")
	http.HandleFunc("/", index)
	http.HandleFunc("/senddata", saveData)
	http.HandleFunc("/form", form)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("index.html")
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := tpl.Execute(w, defaultPerson); err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	return
}

func saveData(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	defaultPerson.Name = req.FormValue("name")
	defaultPerson.Age = req.FormValue("age")
	defaultPerson.Game = req.FormValue("game")
	http.Redirect(w, req, "/", http.StatusFound)
	return
}

func form(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "form.html")
	return
}

type Person struct {
	Name string 
	Age string 
	Game string 
}