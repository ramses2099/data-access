package model

import "time"

type Product struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int64     `json:"stock"`
	Create_at   time.Time `json:"create_at"`
}
