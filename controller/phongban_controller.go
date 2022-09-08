package controller

import (
	"Dichvukhachhang/database"
)

type PhongbanController struct {
	phongbancontroller *database.PhongbanRepo
}

func NewPhongbanController(phongbancontroller *database.PhongbanRepo) *PhongbanController {
	return &PhongbanController{
		phongbancontroller: phongbancontroller,
	}
}
func toPhongbanModel(phongban Phongban) database.PhongbanModel {
	return database.PhongbanModel{
		Name: phongban.Name,
		ID:   phongban.ID,
	}
}
func fromPhongbanModel(model database.PhongbanModel) Phongban {
	return Phongban{
		Name: model.Name,
		ID:   model.ID,
	}
}
func (p *PhongbanController) CreatePhongban(phongban Phongban) error {
	model := toPhongbanModel(phongban)
	err := p.phongbancontroller.CreatePhongbanRepo(model)
	if err != nil {
		return err
	}
	return nil
}
func (p *PhongbanController) DeleteUser(id int64) error {
	err := p.phongbancontroller.DeletePhongbanRepo(id)
	if err != nil {
		return err
	}
	return nil
}
func (p *PhongbanController) UpdatePhongban(phongban Phongban) error {
	model := toPhongbanModel(phongban)
	err := p.phongbancontroller.UpdatePhongbanRepo(model)
	if err != nil {
		return err
	}
	return nil
}

func (p *PhongbanController) GetPhongban(id int64) (phongban Phongban) {
	model, _ := p.phongbancontroller.GetPhongbanRepo(id)
	phongban = fromPhongbanModel(*model)
	return phongban
}
func (p *PhongbanController) GetListPhongban() []Phongban {
	var list []Phongban
	err, _ := p.phongbancontroller.GetListPhongbanRepo()
	for j := 0; j < len(err); j++ {
		list = append(list, fromPhongbanModel(err[j]))
	}
	return list
}
