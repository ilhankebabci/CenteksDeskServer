package model

import "time"

type Department struct {
	Id          int       `xorm:"'id' int notnull pk autoincr"`
	CompanyName string    `xorm:"'company_name' varchar(255) notnull index"`
	Name        string    `xorm:"'name' varchar(255) notnull"`
	CreatedAt   time.Time `xorm:"'created_at' datetime created"`
}

func (m *Department) TableName() string {
	return "department"
}
