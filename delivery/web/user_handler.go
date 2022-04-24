package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/erdauletbatalov/tsarka/domain"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase domain.UserUsecase
}

func NewUserHandler(r *gin.Engine, us domain.UserUsecase) {
	handler := &UserHandler{
		userUsecase: us,
	}

	r.POST("/rest/user", handler.create)
	r.GET("/rest/user/:id", handler.get)
	r.PUT("/rest/user/:id", handler.update)
	r.DELETE("/rest/user/:id", handler.remove)

}

func (u *UserHandler) create(c *gin.Context) {
	user := &domain.User{}

	err := c.BindJSON(user)
	fmt.Println(user, err)
	if err != nil {
		c.Writer.WriteHeader(getStatusCode(err))
		return
	}

	err = u.userUsecase.Create(c.Request.Context(), user)
	if err != nil {
		c.Writer.WriteHeader(getStatusCode(err))
		return
	}

	c.Writer.WriteHeader(http.StatusCreated)
}

func (u *UserHandler) get(c *gin.Context) {
	profileID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || profileID <= 0 {
		c.Writer.WriteHeader(http.StatusBadRequest)
	}

	user, err := u.userUsecase.Get(c.Request.Context(), profileID)
	if err != nil {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (u *UserHandler) update(c *gin.Context) {
	profileID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || profileID <= 0 {
		c.Writer.WriteHeader(http.StatusBadRequest)
	}

	user := &domain.User{
		ID: int64(profileID),
	}

	err = c.BindJSON(user)
	if err != nil {
		c.Writer.WriteHeader(getStatusCode(err))
		return
	}

	err = u.userUsecase.Update(c.Request.Context(), user)
	if err != nil {
		c.Writer.WriteHeader(getStatusCode(err))
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
}

func (u *UserHandler) remove(c *gin.Context) {
	profileID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || profileID <= 0 {
		c.Writer.WriteHeader(http.StatusBadRequest)
	}

	err = u.userUsecase.Delete(c.Request.Context(), profileID)
	if err != nil {
		c.Writer.WriteHeader(getStatusCode(err))
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
}
