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
func (r *Repo) ViewAllJob() ([]models.Job, error) {
	var jobdetials []models.Job
	result := r.Db.Find(&jobdetials)
	if result.Error != nil {
		return []models.Job{}, fmt.Errorf("error in fetching in the jobs :%w", result.Error)
	}
	return jobdetials, nil
}
func (r *Repo) ViewJobById(jid string) (models.Job, error) {
	var jobdetials models.Job
	result := r.Db.Where("id=?", jid).Find(&jobdetials)
	if result.Error != nil {
		return models.Job{}, fmt.Errorf("error in finding job records:%w", result.Error)
	}
	return jobdetials, nil
}
