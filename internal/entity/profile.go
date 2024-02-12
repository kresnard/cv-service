package entity

import "time"

type Profile struct {
	Id         int
	Name       string
	Gender     string
	BirthDate  string
	BirthPlace string
	Email      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
