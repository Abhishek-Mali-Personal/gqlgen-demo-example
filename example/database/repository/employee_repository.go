package repository

import (
	"errors"
	"strings"

	"gorm.io/gorm"

	"example/models"
)

type EmployeeRepository struct {
	DB *gorm.DB
}

func (employeeRepository *EmployeeRepository) CreateEmployee(employee *models.Employee) (*models.Employee, error) {
	db := employeeRepository.DB.Create(&employee)
	err := db.Error
	if err == nil && db.RowsAffected <= 0 {
		err = errors.New("unable to register employee")
	}
	return employee, err
}

func (employeeRepository *EmployeeRepository) CheckEmployeeCountByEmployeeIDs(employeeIDs []int) (int, error) {
	var existCount int64
	db := employeeRepository.DB.Model(&models.Employee{}).Where("id IN ?", employeeIDs).Count(&existCount)
	err := db.Error
	if err == nil && db.RowsAffected <= 0 {
		err = errors.New("unable to find employee")
	}
	return int(existCount), err
}

func (employeeRepository *EmployeeRepository) RetrieveEmployees(filter *models.FilterInput,
	page *int,
) (*models.Pagination, error) {
	var (
		resultQuery *gorm.DB
		employees   []*models.Employee
	)
	pagination := new(models.Pagination)
	pagination.SetPage(page)
	resultQuery = employeeRepository.DB.Model(&models.Employee{}).
		Count(&pagination.TotalRows)
	if filter != nil {
		if filter.EmployeeName != nil && strings.TrimSpace(*filter.EmployeeName) != "" {
			*filter.EmployeeName = "%" + *filter.EmployeeName + "%"
			resultQuery = resultQuery.Where("employees.name ILIKE ?", *filter.EmployeeName).
				Count(&pagination.TotalRows)
		}
		if filter.EmployeeDesignation != nil && strings.TrimSpace(*filter.EmployeeDesignation) != "" {
			*filter.EmployeeDesignation = "%" + *filter.EmployeeDesignation + "%"
			resultQuery = resultQuery.Where("employees.designation ILIKE ?", *filter.EmployeeDesignation).
				Count(&pagination.TotalRows)
		}
		if filter.EmployeeEmail != nil && strings.TrimSpace(*filter.EmployeeEmail) != "" {
			*filter.EmployeeEmail = "%" + *filter.EmployeeEmail + "%"
			resultQuery = resultQuery.Where("employees.email ILIKE ?", *filter.EmployeeEmail).
				Count(&pagination.TotalRows)
		}

		if (filter.TaskName != nil && strings.TrimSpace(*filter.TaskName) != "") ||
			filter.TaskStartDate != nil ||
			filter.TaskEndDate != nil {
			resultQuery = resultQuery.
				Joins("inner join employee_tasks on employee_id = employees.id").
				Joins("inner join tasks on task_id = tasks.id").
				Count(&pagination.TotalRows)
			if filter.TaskName != nil && strings.TrimSpace(*filter.TaskName) != "" {
				*filter.TaskName = "%" + *filter.TaskName + "%"
				resultQuery = resultQuery.Where("tasks.name ILIKE ?", *filter.TaskName).
					Count(&pagination.TotalRows)
			}
			if filter.TaskStartDate != nil {
				resultQuery = resultQuery.Where("tasks.start_date = ?", *filter.TaskStartDate).
					Count(&pagination.TotalRows)
			}
			if filter.TaskEndDate != nil {
				resultQuery = resultQuery.Where("tasks.end_date = ?", *filter.TaskEndDate).
					Count(&pagination.TotalRows)
			}
		}
	}
	resultQuery = resultQuery.Scopes(pagination.Paginate("employees.name asc")).Find(&employees)
	err := resultQuery.Error
	if err == nil && resultQuery.RowsAffected <= 0 {
		return nil, errors.New("no employee found")
	}
	pagination.Rows = models.EmployeeList{
		Employees: employees,
	}
	return pagination, nil
}
