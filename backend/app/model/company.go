package model

import "time"

type Company struct {
	Id        int       `xorm:"'id' int notnull pk autoincr"`
	Name      string    `xorm:"'name' varchar(255) notnull unique"`
	Note      string    `xorm:"'note' varchar(255)"`
	CreatedAt time.Time `xorm:"'created_at' datetime created"`
}

func (m *Company) TableName() string {
	return "company"
}
