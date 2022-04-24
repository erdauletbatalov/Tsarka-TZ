package web

import (
	"net/http"

	"github.com/erdauletbatalov/tsarka/usecase"
	"github.com/gin-gonic/gin"
)

func NewHandler(r *gin.Engine) {
	r.POST("/rest/substr/find", substr)
	r.POST("/rest/email/check", email)
	r.POST("/rest/pin/check", pin)
}

func substr(c *gin.Context) {
	raw, _ := c.GetRawData()
	stringRaw := string(raw)

	// fmt.Println(stringRaw)
	c.JSON(http.StatusOK, gin.H{"find": usecase.LongestSubstring(stringRaw)})
}

func email(c *gin.Context) {
	raw, _ := c.GetRawData()
	stringRaw := string(raw)

	c.JSON(http.StatusOK, gin.H{"find": usecase.Emails(stringRaw)})
}

func pin(c *gin.Context) {
	raw, _ := c.GetRawData()
	stringRaw := string(raw)

	c.JSON(http.StatusOK, gin.H{"find": usecase.PIN(stringRaw)})
}
