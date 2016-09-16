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

// SearchByName : permet de récupérer un utilisateur à partir de son nom
func (service UserService) SearchByName(name string) ([]model.User, error) {
	return service.Dao.GetByName(name)
}

// SearchByFirstName : permet de récupérer un utilisateur à partir de son prénom
func (service UserService) SearchByFirstName(firstName string) ([]model.User, error) {
	return service.Dao.GetByFirstName(firstName)
}

// SearchByPartialFirstNameOrName : permet de récupérer un utilisateur à partir de son nom ou son prénom (en partie)
func (service UserService) SearchByPartialFirstNameOrName(search string) ([]model.User, error) {
	return service.Dao.GetByPartialFirstNameOrName(search)
}

// SearchByEmail : permet de récupérer un utilisateur à partir de son email
func (service UserService) SearchByEmail(email string) ([]model.User, error) {
	return service.Dao.SearchByEmail(email)
}

// GetByEmail : retourne un utilisateur à partir de son email
func (service UserService) GetByEmail(email string) (user model.User, err error) {
	return service.Dao.GetByEmail(email)
}

// GetByPseudo : retour un utilisateur a partir de son pseudo
func (service UserService) GetByPseudo(pseudo string) (user model.User, err error) {
	return service.Dao.GetByPseudo(pseudo)
}

// SearchByPseudo : permet de récupérer un utilisateur à partir de son pseudo
func (service UserService) SearchByPseudo(pseudo string) ([]model.User, error) {
	return service.Dao.SearchByPseudo(pseudo)
}

// UserAuthenticate : authentification d'un utilisateur a partir de son pseudo et son mot de passe
func (service UserService) UserAuthenticate(pseudo string, pwd string) (retour model.User, err error) {
	return service.Dao.Authenticate(pseudo, pwd)
}

// UserAuthenticateByEMail : authentification d'un utilisateur a partir de son email et son mot de passe
func (service UserService) UserAuthenticateByEMail(mail string, pwd string) (retour model.User, err error) {
	return service.Dao.AuthenticateByEmail(mail, pwd)
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
