package checklistcontroller

import (
	"golang-gorm/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetChecklists(c *gin.Context) {
	var checklists []models.Checklist

	models.DB.Preload("ChecklistItems").Find(&checklists)

	c.JSON(http.StatusOK, gin.H{"checklists": checklists})
}

func CreateChecklist(c *gin.Context) {
	var checklist models.Checklist

	if err := c.ShouldBindJSON(&checklist); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&checklist)
	c.JSON(http.StatusOK, gin.H{"checklist": checklist})
}

func DetailChecklist(c *gin.Context) {
	var checklist models.Checklist
	id := c.Param("id")

	if err := models.DB.First(&checklist, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"checklist": checklist})
}

func UpdateChecklist(c *gin.Context) {
	var checklist models.Checklist
	id := c.Param("id")

	if err := c.ShouldBindJSON(&checklist); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&checklist).Where("id = ?", id).Updates(&checklist).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate checklist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})

}

func DeleteChecklist(c *gin.Context) {
	var checklist models.Checklist

	id := c.Param("id")

	if err := models.DB.Where("id = ?", id).First(&checklist).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	models.DB.Delete(&checklist)

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}

func GetItems(c *gin.Context) {
	var checklist []models.ChecklistItem
	id := c.Param("id")

	if err := models.DB.Where("checklist_id = ?", id).Find(&checklist).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"checklist": checklist})
}
