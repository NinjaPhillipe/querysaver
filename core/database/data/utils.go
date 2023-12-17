package data

type Tag struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
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
