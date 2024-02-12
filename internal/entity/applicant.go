package entity

import "time"

type Applicant struct {
	Id          int
	Name        string
	Gender      string
	BirthDate   string
	BirthPlace  string
	Email       string
	Address     string
	City        string
	Province    string
	PostalCode  int
	CompanyName string
	Role        string
	Description string
	Duration    int
	Salary      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
