package repository

import (
	"errors"
	"strings"

	"gorm.io/gorm"

	"example/models"
)

type ComapnyRepository struct {
	DB *gorm.DB
}

func (comapnyRepository *ComapnyRepository) CreateComapny(company *models.Company) (*models.Company, error) {
	db := comapnyRepository.DB.Create(&company)
	err := db.Error
	if err == nil && db.RowsAffected <= 0 {
		err = errors.New("unable to register company")
	}
	return company, err
}

func (comapnyRepository *ComapnyRepository) RetrieveCompanies(filter *models.FilterCompanyInput,
	page *int,
) (*models.Pagination, error) {
	var (
		resultQuery *gorm.DB
		companies   []*models.Company
	)
	pagination := new(models.Pagination)
	resultQuery = comapnyRepository.DB.Model(&models.Company{}).Count(&pagination.TotalRows)
	if filter != nil {
		if filter.CompanyName != nil && strings.TrimSpace(*filter.CompanyName) != "" {
			*filter.CompanyName = "%" + *filter.CompanyName + "%"
			resultQuery = resultQuery.Where("companies.name ILIKE ?", *filter.CompanyName)
		}
		if filter.CompanyWebsite != nil && strings.TrimSpace(*filter.CompanyWebsite) != "" {
			*filter.CompanyWebsite = "%" + *filter.CompanyWebsite + "%"
			resultQuery = resultQuery.Where("companies.website ILIKE ?", *filter.CompanyWebsite)
		}

		if (filter.EmployeeEmail != nil && strings.TrimSpace(*filter.EmployeeEmail) != "") ||
			(filter.EmployeeName != nil && strings.TrimSpace(*filter.EmployeeName) != "") {
			resultQuery = resultQuery.Joins("inner join employees on company_id = companies.id").
				Count(&pagination.TotalRows)
			if filter.EmployeeName != nil && strings.TrimSpace(*filter.EmployeeName) != "" {
				*filter.EmployeeName = "%" + *filter.EmployeeName + "%"
				resultQuery = resultQuery.Where("employees.name ILIKE ?", *filter.EmployeeName)
			}
			if filter.EmployeeEmail != nil && strings.TrimSpace(*filter.EmployeeEmail) != "" {
				*filter.EmployeeEmail = "%" + *filter.EmployeeEmail + "%"
				resultQuery = resultQuery.Where("employees.email ILIKE ?", *filter.EmployeeEmail)
			}
		}
	}
	resultQuery = resultQuery.Scopes(pagination.Paginate("companies.name asc")).Distinct("companies.*").Find(&companies)
	err := resultQuery.Error
	if err == nil && resultQuery.RowsAffected <= 0 {
		return nil, errors.New("no company found")
	}
	pagination.Rows = models.CompanyList{
		Companies: companies,
	}
	return pagination, nil
}
