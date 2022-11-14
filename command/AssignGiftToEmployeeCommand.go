package command

import (
	"errors"
	"fmt"
	"giftem/repo/employeeRepo"
	"giftem/repo/giftRepo"
	"sync"
	"time"
)

type AssignGiftToEmployeeCommand struct {
	mu         sync.Mutex
	giftRepo   giftRepo.GiftsData
	EmployeeId int
}

func NewAssignGiftToEmployeeCommand(employeeId int) *AssignGiftToEmployeeCommand {
	return &AssignGiftToEmployeeCommand{EmployeeId: employeeId}
}

func (c *AssignGiftToEmployeeCommand) Execute() {
	fmt.Println("Executing...")
	employee, err := employeeRepo.FindById(c.EmployeeId)
	if err != nil {
		errors.New("Employee not found.")
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.giftRepo.LoadGifts() // how we could prevent this call?

	foundGift, err := c.giftRepo.FindOneByCategories(employee.Interests)
	if err != nil {
		foundGift = c.giftRepo.FindLast()
	}
	time.Sleep(5 * time.Second)
	fmt.Println(foundGift)
	c.giftRepo.TakeGift(foundGift.Name)
	c.giftRepo.PersistData()
}
