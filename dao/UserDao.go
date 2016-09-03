// Package dao : gestion de l'accès à la base pour la table utilisateur
package dao

import (
	"fmt"
	"log"

	"github.com/tipounet/go-bank/model"
	"gopkg.in/pg.v4"
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
func (dao UserDao) GetByFirstName(firstName string) (user []model.User, err error) {
	err = dao.DB.Model(&user).Where("lower(prenom) like concat('%',lower(?),'%')", firstName).Select()
	return
}

// GetByPartialFirstNameOrName : return user form partial name or first name
func (dao UserDao) GetByPartialFirstNameOrName(search string) (user []model.User, err error) {
	err = dao.DB.Model(&user).Where("lower(prenom) like concat('%',lower(?),'%') or lower(name) like concat('%',lower(?),'%')", search).Select()
	return
}

// SearchByEmail : return user from email
func (dao UserDao) SearchByEmail(email string) (retour []model.User, err error) {
	log.Printf("recherche par mail %v \n", email)
	err = dao.DB.Model(&retour).Select()
	// .Where("lower(email) like concat('%',lower(?),'%')", email)
	// err = dao.DB.Model(&retour).Where("email ='%moogli@phpjungle.info'", email).Select()
	fmt.Printf("les utilisateurs trouvé(s) : %v\n", retour)
	return
}

// SearchByPseudo : return user with pseudo like"%pseudo%"
func (dao UserDao) SearchByPseudo(pseudo string) (user []model.User, err error) {
	err = dao.DB.Model(&user).Where("lower(pseudo) like concat('%',lower(?),'%')", pseudo).Select()
	return
}

//GetByPseudo : retourne un seul utilisateur a partir de son email
func (dao UserDao) GetByPseudo(email string) (user model.User, err error) {
	_, err = dao.DB.QueryOne(&user, "select * from user where pseudo = ?", email)
	return
}

//GetByEmail : retourne un seul utilisateur a partir de son email
func (dao UserDao) GetByEmail(email string) (user model.User, err error) {
	// FIXME : nil dereference, commenton utilise ce QueryOne, ou alors utiliser model pourun select "one" ?
	log.Printf("Recherche d'un utilisateur depuis son email %v\n\n", email)
	defer log.Printf("Utilisateur trouvé %v \n\n", user)
	// _, err = dao.DB.QueryOne(&retour, "select * from user where email = ?", email)
	// us, _ := dao.SearchByEmail(email)
	// user = us[0]
	user = model.User{
		UserID: 42,
		Pseudo: "moogli",
	}
	return
}

// Authenticate Check if pseudo and pwd match
func (dao UserDao) Authenticate(pseudo string, pwd string) (retour bool, err error) {
	dao.DB.Prepare("select count(1) from users where pseudo = ? and pwd = ?")
	// exec + récup + return bool error !
	retour = true
	return
}

// AuthenticateByEmail authentification d'un utilisateur avec son email
func (dao UserDao) AuthenticateByEmail(pseudo string, pwd string) (retour bool, err error) {
	dao.DB.Prepare("select count(1) from users where email = ? and pwd = ?")
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
