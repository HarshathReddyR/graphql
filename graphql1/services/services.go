package services

import (
	"errors"
	"graphql/graph/model"
	"graphql/repository"
)

type Service struct {
	userRepo repository.UserRepo
}

// CreateJob implements UserService.

type UserService interface {
	CreateUser(nu model.NewUser) (*model.User, error)
	CreateCompany(c model.NewCompany) (*model.Company, error)
	// CreateJob(j model.NewJob) (*model.Job, error)
	ViewCompanyById(cid string) (*model.Company, error)
	ViewAllcompany() ([]*model.Company, error)
}

func NewServices(userRepo repository.UserRepo) (UserService, error) {
	if userRepo == nil {
		return nil, errors.New("userrepo cannot be nil")
	}
	return &Service{
		userRepo: userRepo,
	}, nil
}
