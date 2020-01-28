package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-crud/model"
	Model "github.com/golang-crud/model"
)

//FindAll List all todos
func FindAll(c *gin.Context) {
	var u []Model.User
	err := Model.FindAll(&u)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, u)
}

// Create handler user
func Create(c *gin.Context) {
	var u Model.User
	c.Bind(&u)
	err := model.CreateUser(&u)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.JSON(http.StatusCreated, u)
}

// Update usuarios
func Update(c *gin.Context) {
	var u Model.User
	id := c.Params.ByName("id")
	c.Bind(&u)
	err := Model.UpdateUser(&u, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNoContent)
	} else {
		c.JSON(http.StatusOK, nil)
	}
}

// FindOne - Busca de usuario pelo o seu id
func FindOne(c *gin.Context) {
	var u Model.User
	id := c.Params.ByName("id")
	err := Model.FindOne(&u, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNoContent)
	} else {
		c.JSON(http.StatusOK, u)
	}
}
