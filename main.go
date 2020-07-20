package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type IndexPageData struct {
	Name string
	Time string
}

func main() {
	log.Println("Application started successfully")
	//Routes
	http.HandleFunc("/hello", handleIndexPage)

	// blocking call
	// starts the server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Failed to start the app. Error: %s", err)
	}
}

// handler
func handleIndexPage(w http.ResponseWriter, r *http.Request) {
	data := IndexPageData{Name: "Zoman", Time: time.Now().Format(time.Stamp)}

	// check for html template
	html := template.Must(template.ParseFiles("templates/index.html"))

	// get param from request
	name := r.FormValue("name")
	if name != "" {
		log.Printf("request by %s\n", name)
		data.Name = name
	}

	// render template with data items
	err := html.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		log.Println("error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
