package model

// User un utilisateur
type User struct {
	UserID int64  `gorm:"primary_key;column:userid" json:"id"`
	Nom    string `json:""`
	Prenom string `json:"prenom"`
	Pseudo string `json:"pseudo"`
	Email  string `json:"email"`
	Pwd    string `json:"pwd"`
	Pwdbit []byte `gorm:"-" json:"pwdbit"`
}
