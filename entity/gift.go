package entity

import "sync"

type Gift struct {
	Name       string   `json:"name"`
	Categories []string `json:"categories"`
}

type Gifts struct {
	mu    sync.Mutex
	Gifts []Gift
}

func (g Gift) HasCategory(category string) bool {
	for _, c := range g.Categories {
		if c == category {
			return true
		}
	}
	return false
}

func (g Gift) HasCategories(categories []string) bool {
	for _, c := range categories {
		if g.HasCategory(c) == true {
			return true
		}
	}
	return false
}

func (gs *Gifts) RemoveGift(name string) {
	gs.mu.Lock()
	defer gs.mu.Unlock()
	for key, gift := range gs.Gifts {
		if gift.Name == name {
			gs.Gifts = append(gs.Gifts[:key], gs.Gifts[key+1:]...)
			return
		}
	}
}
