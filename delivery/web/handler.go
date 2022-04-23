package web

import (
	"fmt"

	"github.com/erdauletbatalov/tsarka/usecase"
	"github.com/gin-gonic/gin"
)

func NewHandler(r *gin.Engine) {
	r.POST("/rest/substr/find", substr)
	r.POST("/rest/email/check", email)
	r.POST("/rest/pin/check", pin)
	r.POST("/rest/counter/add/", add)
	r.POST("/rest/counter/sub/", sub)
	r.GET("/rest/counter/val", get)
}

func add(c *gin.Context) {
	url := c.Request.URL.String()
	fmt.Println(url)
}

func sub(c *gin.Context) {
	url := c.Request.URL.String()
	fmt.Println(url)
}

func get(c *gin.Context) {
	url := c.Request.URL.String()
	fmt.Println(url)
}

func substr(c *gin.Context) {
	raw, _ := c.GetRawData()
	stringRaw := string(raw)

	// fmt.Println(stringRaw)
	fmt.Println(usecase.LongestSubstring(stringRaw))

}

func email(c *gin.Context) {
	raw, _ := c.GetRawData()
	stringRaw := string(raw)

	fmt.Println(usecase.Emails(stringRaw))
}

func pin(c *gin.Context) {
	raw, _ := c.GetRawData()
	stringRaw := string(raw)

	fmt.Println(usecase.PIN(stringRaw))
}
