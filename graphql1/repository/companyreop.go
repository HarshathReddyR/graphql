package repository

import (
	"errors"
	"fmt"
	"graphql/models"
)

func (r *Repo) CreateCompany(c models.Company) (models.Company, error) {
	result := r.Db.Create(&c)
	if result.Error != nil {
		return models.Company{}, fmt.Errorf("error in create in the creating company the %w", result.Error)
	}
	return c, nil
}
func (r *Repo) ViewCompanyById(cid string) (models.Company, error) {
	var companyDetails models.Company

	result := r.Db.Where("id=?", cid).First(&companyDetails)
	if result.Error != nil {
		return models.Company{}, errors.New("could not found company in the record")
	}
	return companyDetails, nil
}

func (r *Repo) ViewAllCompany() ([]models.Company, error) {
	var companyDetails []models.Company

	result := r.Db.Find(&companyDetails)
	if result.Error != nil {
		return []models.Company{}, errors.New("could not found companies in record")
	}
	return companyDetails, nil
}
