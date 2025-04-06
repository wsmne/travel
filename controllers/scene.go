package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"proj/travel/models"
)

func Add1(c *gin.Context) {
	var req models.Scene
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	scene, err := models.FindSceneByID(req.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Scene not found"})
		return
	}

	if req.Goods == 1 {
		scene.Goods++
	} else if req.Views == 1 {
		scene.Views++
	}
	if _, err := models.UpdateScene(scene); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update scene"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Goods count incremented",
		"goods":   scene.Goods,
	})
}
