package command

import (
	"errors"
	"giftem/entity"
	"giftem/repo/employeeRepo"
	"giftem/repo/giftRepo"
)

type AssignGiftToEmployeeCommand struct {
	giftRepo giftRepo.GiftsData
}

func NewAssignGiftToEmployeeCommand() *AssignGiftToEmployeeCommand {
	return &AssignGiftToEmployeeCommand{}
}

func (c *AssignGiftToEmployeeCommand) Execute(employeeId int) (entity.Gift, error) {
	employee, err := employeeRepo.FindById(employeeId)
	if err != nil {
		errors.New("Employee not found.")
	}
	c.giftRepo.LoadGifts() // how we could prevent this call?

	foundGift, err := c.giftRepo.FindOneByCategories(employee.Interests)
	if err != nil {
		foundGift = c.giftRepo.FindLast()
	}

	c.giftRepo.RemoveGift(foundGift.Name)
	c.giftRepo.PersistData()

	return foundGift, nil
}
