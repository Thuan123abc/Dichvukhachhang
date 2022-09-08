package api

import (
	"Dichvukhachhang/controller"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TicketApi struct {
	ticketcontroller *controller.TicketController
}

func NewTicketApi(Ticketcontroller *controller.TicketController) *TicketApi {
	return &TicketApi{
		ticketcontroller: Ticketcontroller,
	}
}
func (t *TicketApi) Hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello Silve Ticket")
}
func main(ticket controller.Ticket) {

}

type CreateTicketInput struct {
	ID             uint64 `form:"id"`
	IDReception    uint64 `form:"idreception"`
	Phone          string `form:"phone"`
	Comment        string `form:"comment"`
	Content        string `form:"content"`
	Status         bool   `form:"status"`
	TimeCompletion uint64 `form:"timecompletion"`
}

func (t *TicketApi) CreateTicket(c *gin.Context) {
	var input CreateTicketInput
	err := c.BindJSON(&input)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ticket := controller.Ticket{
		ID:             input.ID,
		IDReception:    input.IDReception,
		Phone:          input.Phone,
		Comment:        input.Comment,
		Content:        input.Content,
		Status:         input.Status,
		TimeCompletion: input.TimeCompletion,
	}
	err = t.ticketcontroller.CreateTicket(ticket)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, "OK")
}

type GetTicketOutput struct {
	ticket controller.Ticket `json:"ticket"`
}

func (t *TicketApi) GetTicket(c *gin.Context) {
	ID := c.Param("id")
	IDint, _ := strconv.ParseInt(ID, 10, 64)
	ticket := t.ticketcontroller.GetTicket(IDint)
	c.JSON(http.StatusOK, &GetTicketOutput{
		ticket: ticket,
	})
}
