package product

import "time"

// Model of product
type Model struct {
	ID           uint
	Name         string
	Observations string
	Price        int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Models []*Model

type Storage interface {
	Storage() error
	//Create(*Model) error
	//Update(*Model) error
	//GetAll(*Model) error
	//GetById(uint) error
	//Delete(uint) error
}
