package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/federicoleon/golang-examples/gin_microservice/domain/httperrors"
	"github.com/gin-gonic/gin"
)

// presentation layer

func Get(c *gin.Context) {
	isXml := c.GetHeader("Accept") == "application/xml"
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64) // base, bitsize
	if err != nil {
		httpErr = httperrors.NewBadRequestError("invalid user id")
		respond(c, isXml, httpErr.Code, httpErr)
		return
	}
	user := users[userId]
	if user == nil {
		httpErr = httperrors.NewNotFoundError(fmt.Sprintf("user %d not found", userId))
		respond(c, isXml, httpErr.Code, httpErr)
		return
	}
	respond(c, isXml, http.StatusOK, user)
}

func Create(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		httpErr = httperrors.NewNotFoundError("invalid json body")
		c.JSON(httpErr.Code, httpErr)
		return
	}
	user.Id = currentId
	currentId++

	users[user.Id] = &user
	c.JSON(http.StatusCreated, user)

}
