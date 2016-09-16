// Package dao : gestion de l'accès à la base pour la table utilisateur
package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/tipounet/go-bank/model"
)

// UserDao : accès aux données des utilisateurs
type UserDao struct {
	DB *gorm.DB
}

// Get ça c'est le get de UserDao
func (dao UserDao) Get() (users []model.User, err error) {
	err = dao.DB.Order("nom asc").Find(&users).Error
	return
}

// GetByID : return a user from id
func (dao UserDao) GetByID(id int64) (user model.User, err error) {
	err = dao.DB.First(&user, id).Error
	return
}

// GetByName : return a user from partial name
func (dao UserDao) GetByName(name string) (users []model.User, err error) {
	err = dao.DB.Order("nom asc").Where("lower(nom) like concat('%',lower(?),'%')", name).Find(&users).Error
	return
}

// GetByFirstName : return a user from id
func (dao UserDao) GetByFirstName(firstName string) (users []model.User, err error) {
	dao.DB.Where("lower(prenom) like concat('%',lower(?),'%')", firstName).Find(&users)
	return
}

// GetByPartialFirstNameOrName : return user form partial name or first name
func (dao UserDao) GetByPartialFirstNameOrName(search string) (users []model.User, err error) {
	err = dao.DB.Where("lower(prenom) like concat('%',lower(?),'%') or lower(name) like concat('%',lower(?),'%')", search).Find(&users).Error
	return
}

// SearchByEmail : return user from email
func (dao UserDao) SearchByEmail(email string) (users []model.User, err error) {
	err = dao.DB.Order("name asc").Find(&users).Error
	return
}

// SearchByPseudo : return user with pseudo like"%pseudo%"
func (dao UserDao) SearchByPseudo(pseudo string) (users []model.User, err error) {
	err = dao.DB.Where("lower(pseudo) like concat('%',lower(?),'%')", pseudo).Find(users).Error
	return
}

//GetByPseudo : retourne un seul utilisateur  àpartir de son email
// FIXME : ça pue ?
func (dao UserDao) GetByPseudo(pseudo string) (user model.User, err error) {
	err = dao.DB.Order("name asc").Where("pseudo = ?", pseudo).First(&user).Error
	return
}

//GetByEmail : retourne un seul utilisateur à partir de son email
func (dao UserDao) GetByEmail(email string) (user model.User, err error) {
	err = dao.DB.Where("email = ?", email).First(&user).Error
	return
}

// Authenticate Check if pseudo and pwd match
// FIXME : voir pour juste le Count
func (dao UserDao) Authenticate(pseudo string, pwd string) (retour model.User, err error) {
	retour, err = userAuthRequest(dao, "pseudo = , and pwd = ?", pseudo, pwd)
	return
}

// AuthenticateByEmail authentification d'un utilisateur avec son email
// FIXME idem
func (dao UserDao) AuthenticateByEmail(email string, pwd string) (retour model.User, err error) {
	retour, err = userAuthRequest(dao, "email = ? and pwd = ?", email, pwd)
	return
}

func userAuthRequest(dao UserDao, predicat string, authField string, pwd string) (u model.User, err error) {
	nb := 0
	err = dao.DB.Where(predicat, authField, pwd).First(&u).Count(&nb).Error
	if nb != 1 {
		err = &DAOerror{Code: 001, Message: "Authentication fail : wrong ID or password"}
	}
	return

}

// Create : jesus ?
func (dao UserDao) Create(user *model.User) (e error) {
	e = dao.DB.Set("gorm:save_associations", false).Create(user).Error
	return
}

// Update : osef
func (dao UserDao) Update(user *model.User) (e error) {
	e = dao.DB.Save(user).Error
	return
}

//Delete : suppression d'un utilisateur
func (dao UserDao) Delete(user *model.User) (e error) {
	e = dao.DB.Delete(user).Error
	return
}
