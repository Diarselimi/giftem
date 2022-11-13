package giftRepo

import (
	"encoding/json"
	"errors"
	"fmt"
	"giftem/entity"
	"os"
	"sync"
)

type GiftsData struct {
	mu    sync.Mutex
	gifts []entity.Gift
}

func (gd *GiftsData) FindOneByCategories(categories []string) (entity.Gift, error) {
	gd.mu.Lock()
	defer gd.mu.Unlock()

	fmt.Println(len(gd.gifts))
	for _, gift := range gd.gifts {
		if gift.HasCategories(categories) == true {
			return gift, nil
		}
	}
	return entity.Gift{}, errors.New("No gift found")
}

func (gd *GiftsData) FindLast() entity.Gift {
	gd.mu.Lock()
	defer gd.mu.Unlock()
	return gd.gifts[len(gd.gifts)-1]
}

func (gd *GiftsData) RemoveGift(giftName string) {
	gd.mu.Lock()
	defer gd.mu.Unlock()
	for key, gift := range gd.gifts {
		if gift.Name == giftName {
			gd.gifts = append(gd.gifts[:key], gd.gifts[key+1:]...)
			//fmt.Println(gd.gifts)
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

	fmt.Println("reading-from-file")
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
