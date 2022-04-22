package web

import (
	"fmt"
	"regexp"

	"github.com/gin-gonic/gin"
)

func NewHandler(r *gin.Engine) {
	r.POST("/rest/substr/find", substr)
}

func substr(c *gin.Context) {
	raw, _ := c.GetRawData()
	fmt.Println(string(raw))

	// matched, _ := regexp.MatchString(`(?=(\w+)\1)`, string(raw))
	re := regexp.MustCompile(`{[^{}]*}`)
	fmt.Println(re.FindStringSubmatch(string(raw)))
	fmt.Println(re.FindStringSubmatch("peach punch"))
	fmt.Println(re.FindStringSubmatch("cricket"))

}
