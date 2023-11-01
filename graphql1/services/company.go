package services

import (
	"graphql/graph/model"
	"graphql/models"
	"strconv"
)

func (s *Service) CreateCompany(c model.NewCompany) (*model.Company, error) {
	companydetials := models.Company{
		Name:     c.Name,
		Location: c.Location,
	}
	company, err := s.userRepo.CreateCompany(companydetials)
	if err != nil {
		return nil, err
	}
	cid := strconv.FormatUint(uint64(company.ID), 10)
	return &model.Company{
		Cid:      cid,
		Name:     company.Name,
		Location: company.Location,
	}, nil
}
func (s *Service) ViewCompanyById(cid string) (*model.Company, error) {
	companyDetails, err := s.userRepo.ViewCompanyById(cid)
	if err != nil {
		return nil, err
	}
	cid = strconv.FormatUint(uint64(companyDetails.ID), 10)

	return &model.Company{
		Cid:      cid,
		Name:     companyDetails.Name,
		Location: companyDetails.Location,
	}, nil
}

func (s *Service) ViewAllcompany() ([]*model.Company, error) {
	companyDetails, err := s.userRepo.ViewAllCompany()
	if err != nil {
		return nil, err
	}
	cDatas := []*model.Company{}

	for _, v := range companyDetails {
		cData := model.Company{
			Cid:      strconv.FormatUint(uint64(v.ID), 10),
			Name:     v.Name,
			Location: v.Location,
		}
		cDatas = append(cDatas, &cData)
	}
	return cDatas, nil

}
