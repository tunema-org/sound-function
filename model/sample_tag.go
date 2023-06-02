package model

type Tag struct {
	ID         int
	CategoryID int
	Name       string
}

type Category struct {
	ID   int
	Name string
}
