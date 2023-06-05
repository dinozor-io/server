package gincon

import "github.com/gin-gonic/gin"

type Controller struct {
	Route    *Route
	Response *Response
	Context  *gin.Context
	Callback func(*gin.Context)
}

func (cont *Controller) Init() {
	cont.Response = new(Response)
	cont.Response.Controller = cont
}
