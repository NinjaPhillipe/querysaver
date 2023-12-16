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

type File struct {
	Id       int
	Name     string
	FolderId int
	Label    string
}
