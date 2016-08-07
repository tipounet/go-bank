package service

import (
	"github.com/tipounet/go-bank/dao"
	"github.com/tipounet/go-bank/model"
)

// UserService Service métier de gestion des utilisateurs
type UserService struct {
	Dao *dao.UserDao
}

// Search Recherche d'utilsiateur
func (service UserService) Search(id int64) (model.User, error) {
	return service.Dao.GetByID(id)
}

// Create : création d'un nouvel utilisateur
func (service UserService) Create(user *model.User) error {
	return service.Dao.Create(user)
}

// Read : liste de tout les utilisateurs
func (service UserService) Read() ([]model.User, error) {
	return service.Dao.Get()
}

// Update : mise à jour d'un utilisateurs
func (service UserService) Update(user *model.User) error {
	return service.Dao.Update(user)
}

// Delete : suppression d'un utilisateurs
func (service UserService) Delete(user *model.User) error {
	return service.Dao.Delete(user)
}
