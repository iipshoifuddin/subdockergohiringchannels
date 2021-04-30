// routes/routes.go
package routes

import (
	"gohiringchannels/controllers"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.Static(os.Getenv("PUBLIC_SHARE_IMG"), os.Getenv("PATH_STATIC_ENG"))

	//Use RegisterValidation in the Controllers struct entity
	// and use the key are the registered word
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		v.RegisterValidation("general", bookableGeneral)
		v.RegisterValidation("bookabledate", bookableDate)
		v.RegisterValidation("passwd", bookablePassword)
		v.RegisterValidation("email", validateEmail)
		v.RegisterValidation("usernm", bookableUsername)
	}

	//Login
	r.POST("/login", controllers.Login)
	r.POST("/todo", controllers.CreateTodo)
	r.POST("/refresh", controllers.Refresh)
	r.POST("/logout", controllers.Logout)
	r.POST("/api/v1/engineers/register", controllers.CreateEngineer)

	//Engineers Endpoint EngineerLookEnterpricesDetail
	r.GET("/api/v1/engineers/getAll", controllers.TokenAuthMiddleware(), controllers.FineEngineers)
	r.GET("/api/v1/engineers/find/:id", controllers.TokenAuthMiddleware(), controllers.FineEngineer)
	r.POST("/api/v1/engineers/update/:id", controllers.TokenAuthMiddleware(), controllers.UpdateEngginer)
	r.POST("/api/v1/engineers/uploadimage/:id", controllers.TokenAuthMiddleware(), controllers.UploadShowcaseEngginer)
	r.DELETE("/api/v1/engineers/delete/:id", controllers.TokenAuthMiddleware(), controllers.DeleteAccount)
	r.GET("/api/v1/engineers/lookenterprices", controllers.EngineersLookEnterprices)
	r.GET("/api/v1/engineers/lookenterprice/detail/:id", controllers.EngineerLookEnterpricesDetail)

	//Core Endpoint
	return r
}
