package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Dog struct {
  Name string
  Breed string
}

func main() {
  fmt.Println("Hosting a server on localhost:8080")

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

  http.HandleFunc("/", greet)

  log.Fatal(http.ListenAndServe(":8080", nil))
}
