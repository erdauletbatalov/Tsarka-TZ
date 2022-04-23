package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/erdauletbatalov/tsarka/domain"
	"github.com/gin-gonic/gin"
)

type CounterHandler struct {
	counterUsecase domain.CounterUsecase
}

func NewCounterHandler(r *gin.Engine, counter domain.CounterUsecase) {
	handler := &CounterHandler{
		counterUsecase: counter,
	}

	r.POST("/rest/counter/add/:number", handler.add)
	r.POST("/rest/counter/sub/:number", handler.sub)

	r.GET("/rest/counter/val", handler.get)

}

func (counter *CounterHandler) add(c *gin.Context) {
	param := c.Param("number")
	num, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "field validation failed"})
	}
	err = counter.counterUsecase.Add(c.Request.Context(), num)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "field validation failed"})
		return
	}
	fmt.Println(num)
}

func (counter *CounterHandler) sub(c *gin.Context) {
	param := c.Param("number")
	num, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "field validation failed"})
	}
	err = counter.counterUsecase.Sub(c.Request.Context(), num)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "field validation failed"})
		return
	}
	fmt.Println(num)
}

func (counter *CounterHandler) get(c *gin.Context) {
	num, err := counter.counterUsecase.Get(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "field validation failed"})
		return
	}
	fmt.Println(num)
}
