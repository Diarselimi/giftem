package command

import (
	"errors"
	"fmt"
	"giftem/entity"
	"giftem/repo/employeeRepo"
	"giftem/repo/giftRepo"
	"sync"
)

type AssignGiftToEmployeeCommand struct {
	mu       sync.Mutex
	giftRepo giftRepo.GiftsData
}

func NewAssignGiftToEmployeeCommand() *AssignGiftToEmployeeCommand {
	return &AssignGiftToEmployeeCommand{}
}

func (c *AssignGiftToEmployeeCommand) Execute(employeeId int) (entity.Gift, error) {
	fmt.Println("Executing...")
	employee, err := employeeRepo.FindById(employeeId)
	if err != nil {
		return entity.Gift{}, errors.New("Employee not found.")
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.giftRepo.LoadGifts() // how we could prevent this call?

	foundGift, err := c.giftRepo.FindOneByCategories(employee.Interests)
	if err != nil {
		foundGift = c.giftRepo.FindLast()
	}

	fmt.Println(foundGift)
	c.giftRepo.TakeGift(foundGift.Name)
	c.giftRepo.PersistData()

	return foundGift, nil
}
