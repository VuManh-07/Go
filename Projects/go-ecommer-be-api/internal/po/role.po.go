package po

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID       uint   `gorm:"column:id; type:int; not null; primaryKey; autoIncrement"`
	RoleName string `gorm:"column:role_name;"`
	RoleNote string `gorm:"column:role_note;"`
}

func (r *Role) TableName() string {
	return "go_db_role"
}
