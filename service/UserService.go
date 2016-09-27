package service

import (
	"errors"

	"github.com/tipounet/go-bank/authentication"
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
	retour, err = service.Dao.GetByPseudo(pseudo)
	if err == nil {
		return userAuthenticate(pwd, retour)
	}
	retour = model.User{}
	return
}

// UserAuthenticateByEMail : authentification d'un utilisateur a partir de son email et son mot de passe
func (service UserService) UserAuthenticateByEMail(mail string, pwd string) (retour model.User, err error) {
	retour, err = service.Dao.GetByEmail(mail)
	if err == nil {
		return userAuthenticate(pwd, retour)
	}
	return
}

// Vérifie les informations utilisateur
func userAuthenticate(pwd string, user model.User) (retour model.User, err error) {
	if checkUserPassword(pwd, user) {
		return user, nil
	}
	err = errors.New("ID and password does not match")
	retour = model.User{}
	return
}

// checkUserPassword : vérifie que le mot de passe fournit match celui en base
func checkUserPassword(pwd string, user model.User) bool {
	return authentication.PasswordMatch(pwd, &authentication.Password{
		Hash: user.Pwd,
		Salt: user.Salted,
	})
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
	// salt := "du sel"
	// dk, _ := scrypt.Key([]byte(user.Pwd), []byte(salt), 16384, 8, 1, 32)
	// user.Pwd = string(dk)
	password := authentication.CreatePassword(user.Pwd)
	user.Pwd = password.Hash

	user.Salted = password.Salt
	// Problème le salt n'est pas utf-8 pas d'insertion => voir pour un blob ? (btea ?)
	return service.Dao.Update(user)
}

// Delete : suppression d'un utilisateurs
func (service UserService) Delete(user *model.User) error {
	return service.Dao.Delete(user)
}
