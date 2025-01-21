package models

type Checklist struct {
	Id             int64  `gorm:"primaryKey" json:"id"`
	Name           string `gorm:"type:varchar(100)" json:"name"`
	ChecklistItems []ChecklistItem
	UserId         uint
}
