package controller

import (
	"app/etc/base"
	mid "app/etc/middleware"
	"app/services/index"

	"github.com/gin-gonic/gin"
)

// IndexController Controller Struct, Extended BaseController
type IndexController struct {
	base.Controller
	Group   *gin.RouterGroup
	Read services.IndexServiceRead
	Modify services.IndexServiceModify
}

// IndexControllerInit Initial Path
func IndexControllerInit(g *gin.RouterGroup) {
	ctrl := &IndexController{Group: g,
		Read: services.IndexServiceReadImpl{},
		Modify: services.IndexServiceModifyImpl{}}
	g.GET("/", ctrl.home)
	g.GET("/me", ctrl.me).Use(mid.TokenAuthMiddleware())
	g.GET("/directory", ctrl.directory)
	g.GET("/help", ctrl.help)
}

func (ctrl *IndexController) home(c *gin.Context) {
	res := ctrl.Modify.HomeCreate()
	c.JSON(200, ctrl.ResOk(res))
}

func (ctrl *IndexController) me(c *gin.Context) {
	res := ctrl.Read.Me()
	c.JSON(200, ctrl.ResOk(res))
}

func (ctrl *IndexController) directory(c *gin.Context) {
	res := ctrl.Read.Directory()
	c.JSON(200, ctrl.ResOk(res))
}

func (ctrl *IndexController) help(c *gin.Context) {
	res := ctrl.Read.Help()
	c.JSON(200, ctrl.ResOk(res))
}