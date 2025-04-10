package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func ResponseCreated(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Created successfully",
	})
}

func ResponseCreatedWithObjectId(c *gin.Context, id uuid.UUID) {
	c.JSON(http.StatusCreated, gin.H{
		"success":   true,
		"message":   "Created successfully",
		"object-id": id,
	})
}

func ResponseUpdated(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Updated successfully",
	})
}

func ResponseDeleted(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Deleted successfully",
	})
}

func ResponseGetWithPagination(c *gin.Context, data interface{}, paging interface{}, filters interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
		"paging":  paging,
		"filters": filters,
	})
}
