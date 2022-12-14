package giftRepo

import (
	"encoding/json"
	"errors"
	"fmt"
	"giftem/entity"
	"os"
)

type GiftsData struct {
	gifts []entity.Gift
}

func (gd *GiftsData) FindOneByCategories(categories []string) (entity.Gift, error) {

	for _, gift := range gd.gifts {
		if gift.IsGifted == false && gift.HasCategories(categories) == true {
			return gift, nil
		}
	}
	return entity.Gift{}, errors.New("No gift found")
}

func (gd *GiftsData) FindLast() entity.Gift {

	for _, gift := range gd.gifts {
		if gift.IsGifted == false {
			return gift
		}
	}
	return entity.Gift{}
}

func (gd *GiftsData) TakeGift(giftName string) {

	for key, gift := range gd.gifts {
		if gift.Name == giftName {
			gd.gifts[key].IsGifted = true
			return
		}
	}
	fmt.Println(gd.gifts)
}

func (gd *GiftsData) PersistData() {

	jsonData, err := json.Marshal(gd.gifts)
	if err != nil {
		fmt.Println("Error while saving")
	}

	err = os.WriteFile("repo/gifts.json", jsonData, 0666)
	if err != nil {
		fmt.Println("Error while writing file")
	}
}

func (gd *GiftsData) LoadGifts() {

	fmt.Println(len(gd.gifts))
	if len(gd.gifts) > 0 {
		return
	}

	content, err := os.ReadFile("repo/gifts.json")
	if err != nil {
		fmt.Println("Could not load")
	}

	gd.prepareData(content)
}

func (gd *GiftsData) prepareData(content []byte) {
	var gg []entity.Gift
	err := json.Unmarshal(content, &gg)
	if err != nil {
		fmt.Println("Decoding failed")
	}
	gd.gifts = gg
}
