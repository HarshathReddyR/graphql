package models

import "gorm.io/gorm"

type Job struct {
	gorm.Model
	JobTitle  string  `json:"jobtitle" validate:"required"`
	JobSalary string  `json:"jobsalary" validate:"required"`
	Company   Company `gorm:"ForeignKey:cid"`
	Cid       uint64  `JSON:"cid, omitempty"`
}
