package model

import "strconv"

// User un utilisateur
type User struct {
	UserID int64  `gorm:"primary_key;column:userid" json:"id"`
	Nom    string `json:"nom"`
	Prenom string `json:"prenom"`
	Pseudo string `json:"pseudo"`
	Email  string `json:"email"`
	Pwd    string `json:"pwd"`
	Salt   string `json:"-" gorm:"-"`
	Salted string `json:"-"`
}

func (u User) String() string {
	return "{\n\tID : " + strconv.FormatInt(u.UserID, 10) + "\n\t" +
		"Nom : " + u.Nom + "\n\t" +
		"Prenom : " + u.Prenom + "\n\t" +
		"Pseudo : " + u.Pseudo + "\n\t" +
		"Email : " + u.Email + "\n}"
}
