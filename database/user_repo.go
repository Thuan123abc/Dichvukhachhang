package database

import "gorm.io/gorm"

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	db.AutoMigrate(&UserModel{})
	return &UserRepo{
		db: db,
	}
}
func (i *UserRepo) CreateUserRepo(model UserModel) error {
	err := i.db.Create(&model).Error
	if err != nil {
		return err
	}
	return nil
}
func (i *UserRepo) DeleteUserRepo(id int64) error {
	var model UserModel
	err := i.db.Delete(&model).Where("id=?", id).Error
	if err != nil {
		return err
	}
	return nil
}

func (i *UserRepo) UpdateUserRepo(model UserModel) error {
	err := i.db.Model(&model).Omit("id").Updates(model).Error
	if err != nil {
		return err
	}
	return nil
}
func (i *UserRepo) GetUserRepo(id int64) (*UserModel, error) {
	var model UserModel
	err := i.db.First(model).Where("id=?", id).Error
	if err != nil {
		return nil, err
	}
	return &model, nil
}
func (i *UserRepo) GetListUserRepo() ([]UserModel, error) {
	var list []UserModel
	err := i.db.Find(list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}
