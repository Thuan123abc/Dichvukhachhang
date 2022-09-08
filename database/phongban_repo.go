package database

import "gorm.io/gorm"

type PhongbanRepo struct {
	db *gorm.DB
}

func NewPhongbanRepo(db *gorm.DB) *PhongbanRepo {
	db.AutoMigrate(&PhongbanModel{})
	return &PhongbanRepo{
		db: db,
	}
}
func (p *PhongbanRepo) CreatePhongbanRepo(model PhongbanModel) error {
	err := p.db.Create(&model).Error
	if err != nil {
		return err
	}
	return nil
}
func (p *PhongbanRepo) DeletePhongbanRepo(id int64) error {
	var model PhongbanModel
	err := p.db.Delete(&model).Where("id=?", id).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *PhongbanRepo) UpdatePhongbanRepo(model PhongbanModel) error {
	err := p.db.Model(&model).Omit("id").Updates(model).Error
	if err != nil {
		return err
	}
	return nil
}
func (p *PhongbanRepo) GetPhongbanRepo(id int64) (*PhongbanModel, error) {
	var model PhongbanModel
	err := p.db.First(model).Where("id=?", id).Error
	if err != nil {
		return nil, err
	}
	return &model, nil
}
func (p *PhongbanRepo) GetListPhongbanRepo() ([]PhongbanModel, error) {
	var list []PhongbanModel
	err := p.db.Find(list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}
