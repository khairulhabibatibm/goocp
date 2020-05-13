package models

type Todo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (b *Todo) TableName() string {
	return "todo"
}
