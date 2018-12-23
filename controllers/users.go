package controllers

import (
	"errors"
	"net/http"
	"shade/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	err := models.AddUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func GetUsers(c *gin.Context) {
	var user []models.User
	err := models.GetUsers(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func GetUser(c *gin.Context) {
	id, error := strconv.Atoi(c.Param("id"))
	if error != nil {
		errors.New("Error getting id")
	}
	var user models.User
	err := models.GetUser(&user, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, user)
	}
}
