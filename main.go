package main

import (
	"fmt"
	"giftem/application"
	"giftem/command"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var mutex = sync.Mutex{}

func handler(w http.ResponseWriter, r *http.Request) {
	employeeId, err := strconv.Atoi(r.URL.Path[1:])
	if err != nil {
		fmt.Fprintf(w, "Employee with id %s not found", r.URL.Path[1:])
		return
	}
	mediator := application.CommandMediator{Mu: &mutex}
	mediator.Add(command.NewAssignGiftToEmployeeCommand(employeeId))

	mediator.Run(&w)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
