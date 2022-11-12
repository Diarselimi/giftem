package main

import (
	"fmt"
	"giftem/command"
	"log"
	"net/http"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {

	employeeId, err := strconv.Atoi(r.URL.Path[1:])
	if err != nil {
		fmt.Fprintf(w, "Employee with id %s not found", r.URL.Path[1:])
		return
	}

	gift, err := command.Execute(employeeId)
	if err != nil {
		fmt.Fprintf(w, "We could not find a gift for you.")
	}
	fmt.Fprintf(w, "We have found a gift for you,<h1> %s </h1>", gift.Name)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
