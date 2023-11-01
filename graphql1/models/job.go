package models

import "gorm.io/gorm"

type Job struct {
	gorm.Model
	JobTitle  string  `json:"job_title" validate:"required"`
	JobSalary string  `json:"job_salary" validate:"required"`
	Company   Company `gorm:"ForeignKey:uid"`
	Uid       uint64  `JSON:"uid, omitempty"`
}
type NewJob struct {
	JobTitle  string `json:"job_title" validate:"required"`
	JobSalary string `json:"job_salary" validate:"required"`
}
