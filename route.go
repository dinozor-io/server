package gincon

import (
	"github.com/gin-gonic/gin"
)

type Route struct {
	Method   int8
	Path     string
	Callback func(*Controller)
}

func (r Route) Init() *Controller {
	cont := new(Controller)
	cont.Init()
	cont.Callback = func(c *gin.Context) {
		cont.Context = c
		r.Callback(cont)
		cont.Context.JSON(cont.Response.HttpStatusCode, cont.Response.Data)
	}
	return cont
}
