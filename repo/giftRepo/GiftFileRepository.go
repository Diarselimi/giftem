package giftRepo

import (
	"encoding/json"
	"errors"
	"fmt"
	"giftem/entity"
	"os"
)

func FindOneByCategories(categories []string) (entity.Gift, error) {
	for _, gift := range getGifts() {
		if gift.HasCategories(categories) == true {
			return gift, nil
		}
	}
	return entity.Gift{}, errors.New("No gift found")
}

func FindLast() entity.Gift {
	gifts := getGifts()

	return gifts[len(gifts)-1]
}

func getGifts() []entity.Gift {
	content, err := os.ReadFile("repo/gifts.json")
	if err != nil {
		fmt.Println("Could not load")
	}

	return prepareData(content)
}

func prepareData(content []byte) []entity.Gift {
	var gifts []entity.Gift
	err := json.Unmarshal(content, &gifts)
	if err != nil {
		fmt.Println("Decoding failed")
	}
	return gifts
}
