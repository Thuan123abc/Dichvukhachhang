package controller

import "Dichvukhachhang/database"

type TicketController struct {
	ticketcontroller *database.TicketRepo
}

func NewTicketController(ticketcontroller *database.TicketRepo) *TicketController {
	return &TicketController{
		ticketcontroller: ticketcontroller,
	}
}
func toTicketModel(ticket Ticket) database.TicketModel {
	return database.TicketModel{
		ID:             ticket.ID,
		IDReception:    ticket.IDReception,
		Phone:          ticket.Phone,
		Comment:        ticket.Comment,
		Content:        ticket.Content,
		Status:         ticket.Status,
		TimeCompletion: ticket.TimeCompletion,
	}
}
func fromTicketModel(model database.TicketModel) Ticket {
	return Ticket{
		ID:             model.ID,
		IDReception:    model.IDReception,
		Phone:          model.Phone,
		Comment:        model.Comment,
		Content:        model.Content,
		Status:         model.Status,
		TimeCompletion: model.TimeCompletion,
	}
}
func (t *TicketController) CreateTicket(ticket Ticket) error {
	model := toTicketModel(ticket)
	err := t.ticketcontroller.CreateTicketRepo(model)
	if err != nil {
		return err
	}
	return nil
}
func (t *TicketController) DeleteTicket(id int64) error {
	err := t.ticketcontroller.DeleteTicketRepo(id)
	if err != nil {
		return err
	}
	return nil
}
func (t *TicketController) UpdateTicket(ticket Ticket) error {
	model := toTicketModel(ticket)
	err := t.ticketcontroller.UpdateTicketRepo(model)
	if err != nil {
		return err
	}
	return nil
}

func (t *TicketController) GetTicket(id int64) (ticket Ticket) {
	model, _ := t.ticketcontroller.GetTicketRepo(id)
	ticket = fromTicketModel(*model)
	return ticket
}
func (t *TicketController) GetListTicket() []Ticket {
	var list []Ticket
	err, _ := t.ticketcontroller.GetListTicketRepo()
	for j := 0; j < len(err); j++ {
		list = append(list, fromTicketModel(err[j]))
	}
	return list
}
