package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Defining our "Dog" data schema
type Dog struct {
  Name string
  Breed string
}

// main() function is automatically called and must
// be apart of some package for execution
func main() {
  fmt.Println("Hosting a server on localhost:8080")

  // := denotes assignment
  // We are using Golang's templates and using a dogs object to
  // populate the template with data
  greet := func (write http.ResponseWriter, request *http.Request) {
    tmpl := template.Must(template.ParseFiles("index.html"))
    dogs := map[string][]Dog {
      "Dogs": {
        { Name: "Bear", Breed: "Maltese" },
        { Name: "Lily", Breed: "Chihuahua" },
        { Name: "Coco", Breed: "Corgi" },
        { Name: "Mochi", Breed: "Chow Chow" },
        { Name: "Oliver", Breed: "Golden Doodle" },
        { Name: "Yuki", Breed: "Akita" },
      },
    }
    tmpl.Execute(write, dogs)
  }

  // take form submission and the post values and render to list
  addDog := func (write http.ResponseWriter, request *http.Request) {
    name := request.PostFormValue("name")
    breed := request.PostFormValue("breed")
    htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'>%s - %s</li>", name, breed)
    tmpl, _ := template.New("t").Parse(htmlStr)
    tmpl.Execute(write, nil)
  }

  // greet() will be run when a user loads the root URL
  http.HandleFunc("/", greet)
  // addDog() lives on add-dog and awaits form submission from user
  http.HandleFunc("/add-dog/", addDog)

  // fail-safe default and provide the server
  log.Fatal(http.ListenAndServe(":8080", nil))
}
