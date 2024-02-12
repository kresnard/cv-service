package entity

import "time"

type Addres struct {
	Id         int
	ProfileId  int
	Address    string
	City       string
	Province   string
	PostalCode int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
