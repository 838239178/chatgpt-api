package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthAPI struct{}

func (h HealthAPI) RegisterRoute(r gin.IRouter) {
	r.GET("/ping", h.Ping)
}

func (HealthAPI) Ping(c *gin.Context) {
	c.Status(http.StatusOK)
}
