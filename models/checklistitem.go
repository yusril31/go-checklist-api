package models

type ChecklistItem struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"type:varchar(100)" json:"name"`
	ChecklistId uint
	Status      int64 `gorm:"type:varchar(50)" json:"status"`
}
