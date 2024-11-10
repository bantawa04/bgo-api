package model

type User struct {
	Base
	FirstName      string `gorm:"first_name" json:"first_name" validate:"required"`
	Email          string `gorm:"email" json:"email" validate:"required,email"`
	LastName       string `gorm:"last_name" json:"last_name" validate:"required"`
	Phone          string `gorm:"phone" json:"phone"  validate:"required,phone"`
	Gender         string `gorm:"gender" json:"gender" validate:"required,gender"`
	Password       string `gorm:"password" json:"password" validate:"required"`
	Status         int8   `gorm:"status" json:"status" validate:"required,status"`
	ProfilePicture string `gorm:"profile_picture" json:"profile_picture"`
}

// TableName gives table name of model
func (u *User) TableName() string {
	return "users"
}

// ToMap convert User to map
func (u *User) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"email":          u.Email,
		"firstName":      u.FirstName,
		"lastName":       u.LastName,
		"phone":          u.Phone,
		"gender":         u.Gender,
		"status":         u.Status,
		"profilePicture": u.ProfilePicture,
	}
}
