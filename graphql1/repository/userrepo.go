package repository

import "graphql/models"

func (r *Repo) CreateUser(u models.User) (models.User, error) {
	result := r.Db.Create(&u)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return u, nil
}
