package repository

import (
	"errors"
	"strings"

	"gorm.io/gorm"

	"example/models"
)

type TaskRepository struct {
	DB *gorm.DB
}

func (taskRepository *TaskRepository) CreateTask(task *models.Task) (*models.Task, error) {
	db := taskRepository.DB.Create(&task)
	err := db.Error
	if err == nil && db.RowsAffected <= 0 {
		err = errors.New("unable to create new task")
	}
	return task, err
}

func (taskRepository *TaskRepository) DeleteTask(taskID int) error {
	db := taskRepository.DB.Delete(&models.Task{}, taskID)
	err := db.Error
	if err == nil && db.RowsAffected <= 0 {
		err = errors.New("unable to delete new task")
	}
	return err
}

func (taskRepository *TaskRepository) RetrieveTaskByID(taskID int) (*models.Task, error) {
	var tasks models.Task
	db := taskRepository.DB.Model(&models.Task{}).Where("id = ?", taskID).First(&tasks)
	err := db.Error
	if err == nil && db.RowsAffected <= 0 {
		err = errors.New("unable to get task")
	}
	return &tasks, err
}

func (taskRepository *TaskRepository) RetrieveTasks(filter *models.FilterInput,
	page *int,
) (*models.Pagination, error) {
	var (
		resultQuery *gorm.DB
		tasks       []*models.Task
	)
	pagination := new(models.Pagination)
	pagination.SetPage(page)
	resultQuery = taskRepository.DB.Model(&models.Task{}).
		Count(&pagination.TotalRows)
	if filter != nil {
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
		if (filter.EmployeeName != nil && strings.TrimSpace(*filter.EmployeeName) != "") ||
			(filter.EmployeeDesignation != nil && strings.TrimSpace(*filter.EmployeeDesignation) != "") ||
			(filter.EmployeeEmail != nil && strings.TrimSpace(*filter.EmployeeEmail) != "") {
			resultQuery = resultQuery.
				Joins("inner join employee_tasks on task_id = tasks.id").
				Joins("inner join employees on employee_id = employees.id").
				Count(&pagination.TotalRows)
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
		}
	}
	resultQuery = resultQuery.Scopes(pagination.Paginate("tasks.name asc")).Distinct("tasks.*").Find(&tasks)
	err := resultQuery.Error
	if err == nil && resultQuery.RowsAffected <= 0 {
		return nil, errors.New("no tasks found")
	}
	pagination.Rows = models.TaskList{
		Tasks: tasks,
	}
	return pagination, nil
}
