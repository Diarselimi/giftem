package entity

type Gift struct {
	Name       string   `json:"name"`
	Categories []string `json:"categories"`
}

type Gifts []Gift

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
