package model

import (
	"github.com/google/uuid"
	"time"
)

type Company struct {
	Id          string    `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}

func NewCompany() *Company {
	id := uuid.New().String()
	return &Company{
		Id:        id,
		CreatedAt: time.Now(),
	}
}
