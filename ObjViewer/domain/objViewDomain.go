package domain

import "time"

type Order struct {
	Id      string
	Time    time.Time
	Price   float64
	Success bool
}

type OrderCreate struct {
	Id    string
	Price float64
	Time  time.Time
}

type OrderCompete struct {
	Id string
}
