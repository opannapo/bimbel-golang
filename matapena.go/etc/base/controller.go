package base

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// Controller ...
type Controller struct {
	JSONRes
	JSONError
}

// ResOk ...
func (bc *Controller) ResOk(data interface{}) *JSONRes {
	return &JSONRes{Data: data, Success: true}
}

// ResFailed ...
func (bc *Controller) ResFailed(err string) *JSONRes {
	return &JSONRes{Error: JSONError{Message: err, Code: 500}, Success: false}
}

func (bc *Controller) OutOk(c *gin.Context, data interface{}) {
	c.JSON(200, JSONRes{Data: data, Success: true})
}

// ResFailed ...
func (bc *Controller) OutFailed(c *gin.Context, code int, err string) {
	if strings.ContainsAny(err, "record found not") {
		code = http.StatusNoContent
	}

	c.JSON(200, &JSONRes{Error: JSONError{Message: err, Code: code}, Success: false})
}
