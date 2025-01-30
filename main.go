package main

import (
	"go-checklist-api/controllers/checklistcontroller"
	"go-checklist-api/controllers/checklistitemcontroller"
	"go-checklist-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	// api route
	r.GET("/api/checklists", checklistcontroller.GetChecklists)
	r.POST("/api/checklist", checklistcontroller.CreateChecklist)
	r.GET("/api/checklist/:id", checklistcontroller.DetailChecklist)
	r.PUT("/api/checklist/:id", checklistcontroller.UpdateChecklist)
	r.DELETE("/api/checklist/:id", checklistcontroller.DeleteChecklist)

	r.GET("/api/checklist/:id/item", checklistitemcontroller.GetItems)

	r.Run()
}
