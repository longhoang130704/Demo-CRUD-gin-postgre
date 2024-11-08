package controllers

import (
	"GIN-TUTORIAL/db"
	"GIN-TUTORIAL/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// [GET] /course/
func GetAllCourses(c *gin.Context) {
    title := "Course Page"
    var courses []models.Course
    if err := db.DB.Find(&courses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Failed to create course: %v", err),
		})
		return
	}
    c.HTML(http.StatusOK, "courseList.html", gin.H{
        "title": title,
        "courses": courses,
    })
}

// [GET] /course/create
func ShowCreateCourse(c *gin.Context) {
    title := "Create Course Page"
    
    c.HTML(http.StatusOK, "create.html", gin.H{
        "title": title,
    })
    
}

// [POST] /course/create
func CreateCourse(c *gin.Context) {
	var course models.Course

	if err := c.ShouldBind(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Failed to bind course data: %v", err),
		})
		return
	}

	if err := db.DB.Create(&course).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Failed to create course: %v", err),
		})
		return
	}

	c.Redirect(http.StatusSeeOther, "/course")
}

// [GET] /course/get
func ShowCourse(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Show Course"})
}
