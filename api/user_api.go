package api

import (
	"Dichvukhachhang/controller"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserApi struct {
	usercontroller *controller.UserController
}

func NewUserApi(usercontroller *controller.UserController) *UserApi {
	return &UserApi{
		usercontroller: usercontroller,
	}
}
func (u *UserApi) Hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello Silve")
}

type CreateUserInput struct {
	ID      uint64 `form:"id"`
	Phone   string `form:"phone"`
	Comment string `form:"comment"`
}

func (u *UserApi) CreateUser(c *gin.Context) {
	var input CreateUserInput
	err := c.BindJSON(&input)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	user := controller.User{
		ID:      input.ID,
		Phone:   input.Phone,
		Comment: input.Comment,
	}
	err = u.usercontroller.CreateUser(user)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, "OK")
}

type GetUserOutput struct {
	user controller.User `json:"name"`
}

func (u *UserApi) GetUser(c *gin.Context) {
	ID := c.Param("id")
	IDint, _ := strconv.ParseInt(ID, 10, 64)
	user := u.usercontroller.GetUser(IDint)
	c.JSON(http.StatusOK, &GetUserOutput{
		user: user,
	})
}
