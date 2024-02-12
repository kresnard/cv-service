package entity

import "time"

type Experience struct {
	Id          int
	ProfileId   int
	CompanyName string
	Role        string
	Description string
	Duration    int
	Salary      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
