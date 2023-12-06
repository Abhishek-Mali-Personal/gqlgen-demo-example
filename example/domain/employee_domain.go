package domain

import "example/models"

func (domain *Domain) CreateNewEmployee(input models.CreateEmployee) (*models.Employee, error) {
	return domain.EmployeeRepository.CreateEmployee(&models.Employee{
		Name:        input.Name,
		Designation: input.Designation,
		Email:       input.Email,
		CompanyID:   input.CompanyID,
	})
}

func (domain *Domain) RetrieveEmployees(filter *models.FilterInput, page *int) (*models.Pagination, error) {
	return domain.EmployeeRepository.RetrieveEmployees(filter, page)
}
