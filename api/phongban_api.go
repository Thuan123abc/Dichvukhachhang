package api

import (
	"Dichvukhachhang/controller"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PhongbanApi struct {
	phongbancontroller *controller.PhongbanController
}

func NewPhongbanApi(phongbancontroller *controller.PhongbanController) *PhongbanApi {
	return &PhongbanApi{
		phongbancontroller: phongbancontroller,
	}
}
func (p *PhongbanApi) Hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello Silve PB")
}

type CreatePhongbanInput struct {
	ID   uint64 `form:"id"`
	Name string `form:"name"`
}

func (p *PhongbanApi) CreatePhongban(c *gin.Context) {
	var input CreatePhongbanInput
	err := c.BindJSON(&input)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	phongban := controller.Phongban{
		ID:   input.ID,
		Name: input.Name,
	}
	err = p.phongbancontroller.CreatePhongban(phongban)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, "OK")
}

type GetPhongbanOutput struct {
	phongban controller.Phongban `json:"phongban"`
}

func (p *PhongbanApi) GetPhongban(c *gin.Context) {
	ID := c.Param("id")
	IDint, _ := strconv.ParseInt(ID, 10, 64)
	phongban := p.phongbancontroller.GetPhongban(IDint)
	c.JSON(http.StatusOK, &GetPhongbanOutput{
		phongban: phongban,
	})
}
