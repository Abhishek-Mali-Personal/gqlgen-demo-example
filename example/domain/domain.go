package domain

import "example/database/repository"

type Domain struct {
	TaskRepository     repository.TaskRepository
	EmployeeRepository repository.EmployeeRepository
	M2MRepository      repository.M2MRepository
	ComapnyRepository  repository.ComapnyRepository
}
