package domain

import "example/models"

func (domain *Domain) RegisterNewComapny(input models.CreateCompany) (*models.Company, error) {
	return domain.ComapnyRepository.CreateComapny(&models.Company{
		Name:    input.Name,
		Website: input.Website,
	})
}

func (domain *Domain) RetrieveCompanies(filter *models.FilterCompanyInput, page *int) (*models.Pagination, error) {
	return domain.ComapnyRepository.RetrieveCompanies(filter, page)
}
