package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rpg/models"
)

func FindCharacters(c *gin.Context) {
	var characters []models.Character

	models.DB.Find(&characters)

	c.JSON(http.StatusOK, gin.H{"data": characters})
}

func CreateCharacter(c *gin.Context) {
	var input models.CreateCharacterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	character := models.Character{Name: input.Name, Class: input.Class, Level: 0}
	models.DB.Create(&character)

	c.JSON(http.StatusOK, gin.H{"data": character})
}

func FindCharacter(c *gin.Context) {
	var character models.Character

	if err := models.DB.Where("id = ?", c.Param("id")).First(&character).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Character not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": character})
}

func UpdateCharacter(c *gin.Context) {
	var character models.Character

	if err := models.DB.Where("id = ?", c.Param("id")).First(&character).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Character not found"})
		return
	}

	var input models.UpdateCharacterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&character).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": character})
}

func DeleteCharacter(c *gin.Context) {
	var character models.Character

	if err := models.DB.Where("id = ?", c.Param("id")).First(&character).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Character not found"})
		return
	}

	models.DB.Delete(&character)

	c.JSON(http.StatusOK, gin.H{"data": character}) // Pode retornar bool aqui também, questão de preferencia
}
