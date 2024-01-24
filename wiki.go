package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wrong Page")
}

func Save(field string) {
	data := []byte(field)
	os.WriteFile("test.txt", data, 0777)
}

func loadFile(title string) (*Page, error) {
	filename := title + "test.txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	newContent := r.FormValue("body")
	title := r.URL.Path[len("/view/"):]

	if newContent != "" {
		Save(newContent)
	}

	p, err := loadFile(title)
	if err != nil {
		p = &Page{Title: title}
	}
	t, _ := template.ParseFiles("./html/view.html")
	t.Execute(w, p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadFile(title)
	if err != nil {
		p = &Page{Title: title}
	}
	t, _ := template.ParseFiles("./html/edit.html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
