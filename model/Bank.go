package model

import "strconv"

// Bank : une banque !
type Bank struct {
	BankID int64  `gorm:"primary_key;column:bankid" json:"id"`
	Name   string `json:"name"`
}

// TableName : permet d'indiquer le nom de la table sinon gorm utilise le pluriel (users)
func (Bank) TableName() string {
	return "bank"
}

func (b Bank) String() string {
	return "{\n\tID : " + strconv.FormatInt(b.BankID, 10) + "\n\t" +
		"Nom : " + b.Name + "\n}"
}
