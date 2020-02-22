package internal

type Meals struct {
	Meal []Meal `json:"meals"`
}

type Meal struct {
	Name string `json:"strMeal"`
}

