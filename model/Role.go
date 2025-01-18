package domain

import "time"

type Role struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Role      string    `gorm:"unique,not null" json:"role"`
	User      []User    `gorm:"foreignKey:RoleID" json:"user,omitempty"`
	Slug      string    `gorm:"unique,not null" json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Role) TableName() string {
	return "roles"
}
