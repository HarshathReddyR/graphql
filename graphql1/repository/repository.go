package repository

import (
	"errors"
	"graphql/models"

	"gorm.io/gorm"
)

type Repo struct {
	Db *gorm.DB
}

// CreateUser implements UserRepo.
type UserRepo interface {
	CreateUser(u models.User) (models.User, error)
	CreateCompany(c models.Company) (models.Company, error)
	CreateJob(j models.Job) (models.Job, error)
	ViewCompanyById(cid string) (models.Company, error)
	ViewAllCompany() ([]models.Company, error)
	ViewJobById(jid string) (models.Job, error)
	ViewAllJob() ([]models.Job, error)
}

func NewRepo(db *gorm.DB) (UserRepo, error) {
	if db == nil {
		return nil, errors.New("db should not be nil")
	}
	return &Repo{
		Db: db,
	}, nil
}
