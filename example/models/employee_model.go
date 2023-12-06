package models

type Employee struct {
	ID          int      `gorm:"column:id;AUTO_INCREMENT;primaryKey;" json:"id"`
	Name        string   `gorm:"column:name;type:varchar(255);not null;unique;" json:"name"`
	Designation string   `gorm:"column:designation;type:varchar(255);not null;" json:"designation"`
	Email       string   `gorm:"column:email;type:varchar(255);not null;unique;" json:"email"`
	CompanyID   int      `gorm:"column:company_id;not null;" json:"company_id"`
	Company     *Company `json:"company,omitempty"`
	Tasks       []*Task  `gorm:"many2many:employee_tasks;" json:"tasks,omitempty"`
}

func GetUnassignedEmployees(resultIDs, inputIds []int) []int {
	uniqueIDs := make([]int, 0)
	if len(resultIDs) == len(inputIds) {
		return nil
	}
	if len(resultIDs) == 0 {
		return inputIds
	}
	flag := false
	for i := range inputIds {
		for j := range resultIDs {
			if i == j {
				flag = true
				break
			}
		}
		if !flag {
			uniqueIDs = append(uniqueIDs, inputIds[i])
		}
	}
	return uniqueIDs
}

func (Employee) IsResultUnion() {}
