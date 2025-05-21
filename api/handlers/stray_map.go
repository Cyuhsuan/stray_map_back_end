package handlers

import (
	"net/http"
	"strconv"

	"github.com/Cyuhsuan/stray_map_back_end/internal/service"
	"github.com/gin-gonic/gin"
)

func CreateStrayMap(c *gin.Context) {
	var request service.CreateStrayMapRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service.StrayMapServiceProvider.CreateStrayMap(&request)
}

func GetStrayMapList(c *gin.Context) {
	strayMaps, err := service.StrayMapServiceProvider.GetStrayMapList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": strayMaps})
}

func GetStrayMapDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	strayMap, err := service.StrayMapServiceProvider.GetStrayMapDetail(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": strayMap})
}

func UpdateStrayMap(c *gin.Context) {
	var request service.UpdateStrayMapRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	service.StrayMapServiceProvider.UpdateStrayMap(uint(id), &request)
}

func DeleteStrayMap(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	service.StrayMapServiceProvider.DeleteStrayMap(uint(id))
	c.JSON(http.StatusOK, gin.H{"message": "stray map deleted successfully"})
}
