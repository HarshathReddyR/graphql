package repository

import (
	"fmt"
	"graphql/models"
)

func (r *Repo) CreateJob(j models.Job) (models.Job, error) {
	result := r.Db.Create(&j)
	if result.Error != nil {
		return models.Job{}, fmt.Errorf("error in create in the creating job the %w", result.Error)
	}
	return j, nil
}
