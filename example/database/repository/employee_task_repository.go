package repository

import (
	"errors"

	"gorm.io/gorm"
)

type M2MRepository struct {
	DB *gorm.DB
}

func (m2m *M2MRepository) AssignTaskInBatch(tasks []map[string]any) error {
	db := m2m.DB.Table("employee_tasks").CreateInBatches(tasks, len(tasks))
	err := db.Error
	if err == nil && db.RowsAffected <= 0 {
		err = errors.New("unable to assign task to employees")
	}
	return err
}

func (m2m *M2MRepository) CheckAndRetrieveEmplyeesAssignedToTasks(taskID int, employeeIDs []int) ([]int, error) {
	var employees []int
	db := m2m.DB.Table("employee_tasks").
		Where("task_id = ? AND employee_id IN ?", taskID, employeeIDs).Select("employee_id").Find(&employees)
	err := db.Error
	return employees, err
}
