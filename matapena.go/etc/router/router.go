package router

import (
	"app/controller"
	"app/etc/base"
	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// InitRoute Initial Router / Grouping & Controller
func InitRoute(db *gorm.DB) {
	gin.SetMode(viper.GetString("mode"))
	router := gin.Default()

	//404
	router.NoRoute(func(c *gin.Context) {
		err := base.JSONError{Message: "Page Not Found", Code: 404}
		c.JSON(200, base.JSONRes{
			Success: false,
			Data:    nil,
			Error:   err,
		})
	})

	//init all controllers
	v1 := router.Group("api/v1")
	controller.IndexControllerInit(v1)
	controller.EmployeeControllerInit(v1, db)

	_ = router.Run(viper.GetString(`server.address`))
}
