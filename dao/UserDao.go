// Package dao : gestion de l'accès à la base pour la table utilisateur
package dao

import (
	"gopkg.in/pg.v4"
	"github.com/tipounet/go-bank/model"
)

// UserDao : accès aux données des utilisateurs
type UserDao struct {
	DB *pg.DB
}

// Get ça c'est le get de UserDao
func (dao UserDao) Get() (retour []model.User, err error) {
	err = dao.DB.Model(&retour).Select()
	return
}

// GetByID : return a user from id
func (dao UserDao) GetByID(id int64) (user model.User, err error) {
	err = dao.DB.Model(&user).Where("userid=?", id).Select()
	return
}

// GetByName : return a user from partial name
func (dao UserDao) GetByName(name string) (user []model.User, err error) {
	err = dao.DB.Model(&user).Where("lower(nom) like concat('%',lower(?),'%')", name).Select()
	return
}

// GetByFirstName : return a user from id
func (dao UserDao) GetByFirstName(prenom string) (user []model.User, err error) {
	err = dao.DB.Model(&user).Where("lower(prenom) like concat('%',lower(?),'%')", prenom).Select()
	return
}

// GetByPartialFirstNameOrName : return user form partial name or first name
func (dao UserDao) GetByPartialFirstNameOrName(search string) (user []model.User, err error) {
	err = dao.DB.Model(&user).Where("lower(prenom) like concat('%',lower(?),'%') or lower(name) like concat('%',lower(?),'%')", search).Select()
	return
}

// GetByMail : return user from email
func (dao UserDao) GetByMail(mail string) (user []model.User, err error) {
	err = dao.DB.Model(&user).Where("mail = ?", mail).Select()
	return
}

// GetByPseudo : return user with pseudo "pseudo"
func (dao UserDao) GetByPseudo(pseudo string) (user []model.User, err error) {
	err = dao.DB.Model(&user).Where("pseudo = ?", pseudo).Select()
	return
}

// Authenticate Check if pseudo and pwd match
func (dao UserDao) Authenticate(pseudo string, pwd string) (retour bool, err error) {
	dao.DB.Prepare("select count(1) from users where pseudo = ? and pwd = ?")
	// exec + récup + return bool error !
	retour = true
	return
}

// Create : jesus ?
func (dao UserDao) Create(user *model.User) error {
	return dao.DB.Create(user)
}

// Update : osef
func (dao UserDao) Update(user *model.User) error {
	return dao.DB.Update(user)
}

//Delete : suppression d'un utilisateur
func (dao UserDao) Delete(user *model.User) error {
	return dao.DB.Delete(user)
}
