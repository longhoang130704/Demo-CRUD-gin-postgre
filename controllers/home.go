package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) { 
	c.HTML(http.StatusOK, "courseList.html",gin.H{
		"title": "Index Page",
	})
}