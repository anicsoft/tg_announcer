package model

type Category struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

func NewCategory() *Category {
	return &Category{}
}
