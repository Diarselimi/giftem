package entity

type Employee struct {
	Id        int      `json:"id"`
	Name      string   `json:"name"`
	Interests []string `json:"interests"`
}

func (e Employee) hasInterests(categories []string) bool {
	for _, c := range categories {
		if e.hasInterest(c) == true {
			return true
		}
	}
	return false
}

func (e Employee) hasInterest(category string) bool {
	for _, i := range e.Interests {
		if i == category {
			return true
		}
	}
	return false
}
