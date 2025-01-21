package models

type User struct {
	Id         int64  `gorm:"primaryKey" json:"id"`
	Username   string `gorm:"type:varchar(20)" json:"username"`
	Email      string `gorm:"type:varchar(80)" json:"email"`
	Password   string
	Checklists []Checklist
}
