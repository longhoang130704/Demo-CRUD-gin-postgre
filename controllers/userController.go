package controllers

import (
	"GIN-TUTORIAL/db"
	"GIN-TUTORIAL/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// [POST] /user/create
func CreateUser(c *gin.Context)  {
	name := c.PostForm("nameInput")
	email := c.PostForm("emailInput")
	password := c.PostForm("passwordInput")
	user := models.Users{
		Name:  name,
		Email: email,
		Password: password,
	}
	
	result := db.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error in Create User",
		})
		return 
	} 
	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "Create User successfully",
	// 	"user": user,
	// })
	c.Redirect(http.StatusSeeOther, "account/manage")
}

// [GET] /user/signUp
func SignUp(c *gin.Context)  {
	title := "Sign Up"
	c.HTML(http.StatusOK, "signUp.html", gin.H{
		"title": title,
	})
}

// [DELETE] /user/account/delete/:id
func DeleteUser(c *gin.Context)  {
	fmt.Println("Delete")
	idDeleteUser := c.Param("id")
	fmt.Println(idDeleteUser)
	var user models.Users
    if err := db.DB.First(&user, idDeleteUser).Error; err != nil {
        fmt.Println("Not foung user to delete")
        return
    }
	if err := db.DB.Delete(&user).Error; err != nil {
        fmt.Println("Could not delete user")
        return
    }
}

// [PUT] /user/account/update/:id
// show update Field
func ShowUpdateUser(c *gin.Context) {
	idUpdateUser := c.Param("id")
	var user models.Users
	isError := db.DB.First(&user, idUpdateUser).Error
	if isError != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Not Found User",
		})
		return
	}
	c.HTML(http.StatusOK, "update.html", gin.H{
		"title": "Update Account Page",
		"user": user,
	})
}

// [POST] /user/account/update/:id
func UpdateUser(c *gin.Context)  {
	idUpdateUser := c.Param("id")

    var user models.Users
    if err := db.DB.First(&user, idUpdateUser).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    name := c.PostForm("nameInputUpdate")
    email := c.PostForm("emailInputUpdate")
    password := c.PostForm("passwordInputUpdate")

    user.Name = name
    user.Email = email
    user.Password = password 

    if err := db.DB.Save(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update user"})
        return
    }

    // c.JSON(http.StatusOK, gin.H{
    //     "message": "User updated successfully",
    //     "user":    user,
    // })

	c.Redirect(http.StatusSeeOther, "/user/account/manage")
	
}

// [GET] user/account/manage
func ManageAccount(c *gin.Context) {
	var usersData []models.Users
	db.DB.Find(&usersData)
	// fmt.Println(usersData)
	c.HTML(http.StatusOK, "manageAccount.html", gin.H{
		"title": "Account Page",
		"users": usersData,
	})
}

// [GET] user/account/trash
func ShowTrash(c *gin.Context)  {
	title := "Trash Page"
	var users []models.Users
	db.DB.Unscoped().Where("deleted_at IS NOT NULL").Find(&users)
	c.HTML(http.StatusOK, "trash.html", gin.H{
		"title": title,
		"users": users,
	})
}

// [GET] user/account/restore/:id
func RestoreUser(c *gin.Context)  {
	idRestoreUser := c.Param("id")
	fmt.Println(idRestoreUser)
	var user models.Users

	if isExist := db.DB.Unscoped().First(&user, idRestoreUser).Error; isExist != nil {
		fmt.Println("Not Found Id user is needed Restore")
		return
	}
	user.DeletedAt =  gorm.DeletedAt{}
	if isError := db.DB.Save(&user).Error; isError != nil {
		fmt.Println("Can Restore user")
		return
	}
	c.Redirect(http.StatusSeeOther, "/user/account/trash")
}