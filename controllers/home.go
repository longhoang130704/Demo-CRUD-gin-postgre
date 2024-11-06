package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.Redirect(http.StatusSeeOther, "/user/signUp")
}