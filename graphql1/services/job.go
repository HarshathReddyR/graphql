package services

import (
	"fmt"
	"graphql/graph/model"
	"graphql/models"
	"strconv"
)

func (s *Service) CreateJob(j model.NewJob) (*model.Job, error) {
	id, err := strconv.ParseUint(j.Cid, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("parssing error:%w", err)
	}
	jobdetials := models.Job{
		JobTitle:  j.JobTitle,
		JobSalary: j.JobSalary,
		Cid:       id,
	}
	cj, err := s.userRepo.CreateJob(jobdetials)
	if err != nil {
		return nil, fmt.Errorf("problem in creating the job:%w", err)
	}
	strcid := strconv.FormatUint(cj.Cid, 10)
	strjid := strconv.FormatUint(uint64(cj.ID), 10)
	commpany, err := s.ViewCompanyById(strcid)
	if err != nil {
		return nil, fmt.Errorf("problem in getting company:%w", err)
	}
	return &model.Job{
		ID:        strjid,
		JobTitle:  cj.JobTitle,
		JobSalary: cj.JobSalary,
		Company:   commpany,
	}, err
}
func (s *Service) ViewJobById(jid string) (*model.Job, error) {
	jobdetials, err := s.userRepo.ViewJobById(jid)
	if err != nil {
		return nil, err
	}
	jid = strconv.FormatUint(uint64(jobdetials.ID), 10)
	compid := strconv.FormatUint(uint64(jobdetials.Company.ID), 10)
	company := model.Company{
		Cid:      compid,
		Name:     jobdetials.Company.Name,
		Location: jobdetials.Company.Location,
	}
	return &model.Job{
		ID:        jid,
		JobTitle:  jobdetials.JobTitle,
		JobSalary: jobdetials.JobSalary,
		Company:   &company,
	}, nil
}
func (s *Service) ViewAllJob() ([]*model.Job, error) {
	jobdetials, err := s.userRepo.ViewAllJob()
	if err != nil {
		return nil, err
	}
	jobdatas := []*model.Job{}
	for _, v := range jobdetials {
		jobdata := model.Job{
			ID:        strconv.FormatUint(uint64(v.ID), 11),
			JobTitle:  v.JobTitle,
			JobSalary: v.JobSalary,
			Company: &model.Company{
				Cid: strconv.FormatUint(uint64(v.Cid), 10),
			},
		}
		jobdatas = append(jobdatas, &jobdata)
	}
	return jobdatas, nil
}
