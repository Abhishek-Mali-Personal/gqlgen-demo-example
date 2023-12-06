package models

type Company struct {
	ID        int         `gorm:"column:id;AUTO_INCREMENT;primaryKey;" json:"id"`
	Name      string      `gorm:"column:name;type:varchar(255);not null;unique;" json:"name"`
	Website   *string     `gorm:"column:website;type:varchar(255);unique;" json:"website,omitempty"`
	Employees []*Employee `gorm:"foreignKey:CompanyID;references:ID;constraint:OnDelete:CASCADE;" json:"employees"`
}

func (Company) IsResultUnion() {}
