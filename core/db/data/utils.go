package data

type Tag struct {
	Id    int
	Name  string
	Color string
}

type TagCreate struct {
	Name  string
	Color string
}
