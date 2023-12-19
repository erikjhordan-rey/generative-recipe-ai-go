package models

type Recipe struct {
	Name          string   `json:"name,omitempty"`
	Description   string   `json:"description,omitempty"`
	Ingredients   []string `json:"ingredients,omitempty"`
	TypeOfCuisine string   `json:"type of cuisine,omitempty"`
	Vegetarian    string   `json:"vegetarian,omitempty"`
}
