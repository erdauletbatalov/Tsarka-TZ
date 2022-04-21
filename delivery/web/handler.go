package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHandler(r *gin.Engine) {
	r.POST("/rest/substr/find", substr)
}

func substr(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
