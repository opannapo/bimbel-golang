package controller

import (
	"app/etc/base"
	"app/etc/param"
	services "app/services/employee"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

// EmployeeController /
type EmployeeController struct {
	base.Controller
	Group  *gin.RouterGroup
	Read   services.EmployeeServiceRead
	Modify services.EmployeeServiceModify
}

// EmployeeControllerInit .
func EmployeeControllerInit(g *gin.RouterGroup, db *gorm.DB) {
	ctrl := EmployeeController{Group: g,
		Read:   services.InstanceServiceRead(db),
		Modify: services.InstanceServiceModify(db)}
	g.POST("/employee/create/admin", ctrl.createAdmin)
	g.POST("/employee/create/teacher", ctrl.createTeacher)
	g.POST("/employee/create/general", ctrl.createGeneral)
	g.GET("/employee/get/admins", ctrl.getAdminAll)
	g.GET("/employee/get/admins/:id", ctrl.getAdminById)

}

func (ctrl *EmployeeController) createAdmin(c *gin.Context) {
	println("createAdmin ")

	var p param.AdminCreate
	err := c.ShouldBindJSON(&p)
	if err != nil {
		ctrl.OutFailed(c, http.StatusBadRequest, "Invalid Parameter Request ! "+err.Error())
		return
	}

	result, err := ctrl.Modify.CreateAdmin(p)
	if err != nil {
		ctrl.OutFailed(c, http.StatusInternalServerError, err.Error())
		return
	}
	ctrl.OutOk(c, result)
}

func (ctrl *EmployeeController) createTeacher(c *gin.Context) {

}

func (ctrl *EmployeeController) createGeneral(c *gin.Context) {

}

func (ctrl *EmployeeController) getAdminAll(c *gin.Context) {
	result, err := ctrl.Read.GetAdminAll()
	if err != nil {
		ctrl.OutFailed(c, http.StatusInternalServerError, err.Error())
		return
	}
	ctrl.OutOk(c, result)
}

func (ctrl *EmployeeController) getAdminById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	result, err := ctrl.Read.GetAdminById(id)
	if err != nil {
		ctrl.OutFailed(c, http.StatusInternalServerError, err.Error())
		return
	}
	ctrl.OutOk(c, result)
}
