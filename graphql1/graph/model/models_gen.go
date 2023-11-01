// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Company struct {
	Cid      string `json:"cid"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

type Job struct {
	ID        string   `json:"id"`
	JobTitle  string   `json:"jobTitle"`
	JobSalary string   `json:"jobSalary"`
	Company   *Company `json:"company"`
}

type NewCompany struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

type NewJob struct {
	JobTitle  string `json:"jobTitle"`
	JobSalary string `json:"jobSalary"`
	ID        string `json:"id"`
}

type NewUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
