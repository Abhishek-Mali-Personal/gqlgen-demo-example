package domain

import (
	"errors"

	"example/models"
)

func (domain *Domain) CreateNewTask(input *models.CreateTask) (*models.Task, error) {
	employeeCount, err := domain.EmployeeRepository.CheckEmployeeCountByEmployeeIDs(input.EmployeeIds)
	if err != nil {
		return nil, err
	}
	if employeeCount != len(input.EmployeeIds) {
		return nil, errors.New("unable to find employees")
	}
	tasks, err := domain.TaskRepository.CreateTask(&models.Task{
		Name:        input.Name,
		Description: input.Description,
		StartDate:   input.StartDate,
		EndDate:     input.EndDate,
	})
	taskEmployee := make([]map[string]any, 0)
	for i := range input.EmployeeIds {
		taskEmployee = append(taskEmployee, map[string]any{
			"task_id":     tasks.ID,
			"employee_id": input.EmployeeIds[i],
		})
	}
	err = domain.M2MRepository.AssignTaskInBatch(taskEmployee)
	if err != nil {
		if domain.TaskRepository.DeleteTask(tasks.ID) != nil {
			return nil, err
		}
		return nil, errors.New("unable to create tasks")
	}
	return tasks, nil
}

func (domain *Domain) AssigntaskToEmployees(taskID int, employeeIDs []int) (*models.Task, error) {
	employeeCount, err := domain.EmployeeRepository.CheckEmployeeCountByEmployeeIDs(employeeIDs)
	if err != nil {
		return nil, err
	}
	if employeeCount != len(employeeIDs) {
		return nil, errors.New("unable to find employees")
	}
	empIDs, err := domain.M2MRepository.CheckAndRetrieveEmplyeesAssignedToTasks(taskID, employeeIDs)
	if err != nil {
		return nil, err
	}
	unassignedEmployeeIDs := models.GetUnassignedEmployees(empIDs, employeeIDs)
	taskEmployee := make([]map[string]any, 0)
	for i := range unassignedEmployeeIDs {
		taskEmployee = append(taskEmployee, map[string]any{
			"task_id":     taskID,
			"employee_id": unassignedEmployeeIDs[i],
		})
	}
	err = domain.M2MRepository.AssignTaskInBatch(taskEmployee)
	if err != nil {
		return nil, err
	}
	task, err := domain.TaskRepository.RetrieveTaskByID(taskID)
	if err != nil {
		return nil, errors.New("something went wrong retrieving tasks")
	}
	return task, nil
}

func (domain *Domain) RetrieveTasks(filter *models.FilterInput, page *int) (*models.Pagination, error) {
	return domain.TaskRepository.RetrieveTasks(filter, page)
}
