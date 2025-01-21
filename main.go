package main

import (
	"golang-gorm/controllers/checklistcontroller"
	"golang-gorm/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/checklists", checklistcontroller.GetChecklists)
	r.POST("/api/checklist", checklistcontroller.CreateChecklist)
	r.GET("/api/checklist/:id", checklistcontroller.DetailChecklist)
	r.PUT("/api/checklist/:id", checklistcontroller.UpdateChecklist)
	r.DELETE("/api/checklist/:id", checklistcontroller.DeleteChecklist)

	r.GET("/api/checklist/:id/item", checklistcontroller.GetItems)

	r.Run()
}
