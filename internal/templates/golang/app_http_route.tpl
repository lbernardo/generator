package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRoutes(
	r *gin.Engine,

) {
	group := r.Group("/{{.name}}")
	group.GET("/", func (ctx *gin.Context) {
	    ctx.Status(http.StatusOK)
	})
}
