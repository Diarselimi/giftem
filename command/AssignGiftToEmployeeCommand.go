package command

import (
	"errors"
	"giftem/entity"
	"giftem/repo/employeeRepo"
	"giftem/repo/giftRepo"
)

func Execute(employeeId int) (entity.Gift, error) {
	employee, err := employeeRepo.FindById(employeeId)
	if err != nil {
		errors.New("Employee not found.")
	}

	foundGift, err := giftRepo.FindOneByCategories(employee.Interests)
	if err != nil {
		foundGift = giftRepo.FindLast()
	}

	//giftRepo.removeGift(foundGift) // removes it from repo

	return foundGift, nil
}
