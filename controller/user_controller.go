package controller

import (
	"Dichvukhachhang/database"
)

type UserController struct {
	usercontroller *database.UserRepo
}

func NewUserController(usercontroller *database.UserRepo) *UserController {
	return &UserController{
		usercontroller: usercontroller,
	}
}
func toUserModel(user User) database.UserModel {
	return database.UserModel{
		ID:      user.ID,
		Phone:   user.Phone,
		Comment: user.Comment,
	}
}
func fromUserModel(model database.UserModel) User {
	return User{
		ID:      model.ID,
		Phone:   model.Phone,
		Comment: model.Comment,
	}
}
func (u *UserController) CreateUser(user User) error {
	model := toUserModel(user)
	err := u.usercontroller.CreateUserRepo(model)
	if err != nil {
		return err
	}
	return nil
}
func (u *UserController) DeleteUser(id int64) error {
	err := u.usercontroller.DeleteUserRepo(id)
	if err != nil {
		return err
	}
	return nil
}
func (u *UserController) UpdateUser(user User) error {
	model := toUserModel(user)
	err := u.usercontroller.UpdateUserRepo(model)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserController) GetUser(id int64) (user User) {
	model, _ := u.usercontroller.GetUserRepo(id)
	user = fromUserModel(*model)
	return user
}
func (u *UserController) GetListUser() []User {
	var list []User
	err, _ := u.usercontroller.GetListUserRepo()
	for j := 0; j < len(err); j++ {
		list = append(list, fromUserModel(err[j]))
	}
	return list
}
