package services

import (
	"graphql/graph/model"
	"graphql/models"
	"graphql/pkg"
	"strconv"
)

func (s *Service) CreateUser(nu model.NewUser) (*model.User, error) {
	hashPasswor, err := pkg.HashPassword(nu.Password)
	if err != nil {
		return nil, err
	}
	userDetials := models.User{
		UserName:     nu.Name,
		Email:        nu.Email,
		HashPassword: hashPasswor,
	}
	userDetials, err = s.userRepo.CreateUser(userDetials)
	if err != nil {
		return nil, err
	}
	uid := strconv.FormatUint(uint64(userDetials.ID), 10)
	return &model.User{
		ID:       uid,
		Name:     userDetials.UserName,
		Email:    userDetials.Email,
		Password: userDetials.HashPassword,
	}, nil
}
