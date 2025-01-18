package domain

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"unique,not null" json:"username"`
	Email     string    `gorm:"unique,not null" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	RoleID    uint      `gorm:"not null" json:"role_id"`
	Role      Role      `gorm:"foreignKey:RoleID" json:"role,omitempty"`
	IsMember  bool      `gorm:"not null,default:false" json:"is_member"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}
