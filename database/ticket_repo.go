package database

import "gorm.io/gorm"

type TicketRepo struct {
	db *gorm.DB
}

func NewTicketRepo(db *gorm.DB) *TicketRepo {
	db.AutoMigrate(&TicketModel{})
	return &TicketRepo{
		db: db,
	}
}
func (t *TicketRepo) CreateTicketRepo(model TicketModel) error {
	err := t.db.Create(&model).Error
	if err != nil {
		return err
	}
	return nil
}
func (t *TicketRepo) DeleteTicketRepo(id int64) error {
	var model UserModel
	err := t.db.Delete(&model).Where("id=?", id).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *TicketRepo) UpdateTicketRepo(model TicketModel) error {
	err := t.db.Model(&model).Omit("id").Updates(model).Error
	if err != nil {
		return err
	}
	return nil
}
func (t *TicketRepo) GetTicketRepo(id int64) (*TicketModel, error) {
	var model TicketModel
	err := t.db.First(model).Where("id=?", id).Error
	if err != nil {
		return nil, err
	}
	return &model, nil
}
func (t *TicketRepo) GetListTicketRepo() ([]TicketModel, error) {
	var list []TicketModel
	err := t.db.Find(list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}
