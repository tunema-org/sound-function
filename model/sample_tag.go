package model

type Sample_tag struct {
	Sample_ID int
	Tag_ID    int
}

type Tag struct {
	ID         int
	CategoryID int
	Name       string
}

type Category struct {
	ID   int
	Name string
}
