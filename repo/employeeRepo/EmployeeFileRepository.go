package employeeRepo

import (
	"encoding/json"
	"errors"
	"fmt"
	"giftem/entity"
	"os"
)

func FindById(id int) (entity.Employee, error) {
	for _, e := range getEmployees() {
		if e.Id == id {
			return e, nil
		}
	}

	return entity.Employee{}, errors.New("Employee not found")
}

func getEmployees() []entity.Employee {
	content, err := os.ReadFile("repo/employees.json")
	if err != nil {
		fmt.Println("Could not load")
	}

	return prepareData(content)
}

func prepareData(content []byte) []entity.Employee {
	var employees []entity.Employee
	err := json.Unmarshal(content, &employees)
	if err != nil {
		fmt.Println("Decoding failed")
	}
	return employees
}
