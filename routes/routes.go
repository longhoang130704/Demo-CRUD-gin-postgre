package routes

import (
	"text/template"

	"github.com/gin-gonic/gin"

	"GIN-TUTORIAL/controllers"
)
	

func InitRoutes() *gin.Engine{
	router := gin.Default()
    router.SetFuncMap(template.FuncMap{})
    router.LoadHTMLGlob("templates/*/*")

    router.GET("/", controllers.Home)

    UserRouter := router.Group("/user") 
    {
        UserRouter.GET("/account/update/:id", controllers.ShowUpdateUser) // khong cai them thu vien nhung cho nay la [PUT]
        UserRouter.POST("/account/update/:id", controllers.UpdateUser) 
        UserRouter.DELETE("/account/delete/:id", controllers.DeleteUser)
        UserRouter.GET("/account/restore/:id", controllers.RestoreUser)
        UserRouter.GET("/account/manage", controllers.ManageAccount)
        UserRouter.GET("/account/trash", controllers.ShowTrash)
        UserRouter.GET("/signUp", controllers.SignUp)
        UserRouter.POST("/create", controllers.CreateUser)
    }
    
    CourseRouter := router.Group("/course") 
    {
        CourseRouter.GET("/get", controllers.ShowCourse)
        CourseRouter.POST("/create", controllers.CreateCourse)
        CourseRouter.GET("/create", controllers.ShowCreateCourse)
        CourseRouter.GET("/", controllers.GetAllCourses)
    }

    return router
}