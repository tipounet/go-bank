package model

import "strconv"

// Account : c'est un compte bancaire a quelqu'un
type Account struct {
	BankaccountID int64  `gorm:"primary_key;column:bankaccountid" json:"id"`
	Accountnumber string `json:"number"`
	// Has one does not work
	UserID int64 `gorm:"column:userid" json:"userid"`
	User   User  `json:"user" gorm:"ForeignKey:UserID;AssociationForeignKey:UserID"`
	BankID int64 `gorm:"column:bankid" json:"bankid"`
	Bank   Bank  `json:"bank" gorm:"ForeignKey:BankID;AssociationForeignKey:BankID"`
}

// TableName : permet d'indiquer le nom de la table sinon gorm utilise le pluriel (users)
func (Account) TableName() string {
	return "bankaccount"
}

func (a Account) String() string {
	s := "Account {\n\t id : " + strconv.FormatInt(a.BankaccountID, 10) + " \n\t numero : " + a.Accountnumber + " \n\t userid : " + strconv.FormatInt(a.UserID, 10) + " \n\t bankid : " + strconv.FormatInt(a.BankID, 10) +
		"\n\t User : " + a.User.String() + "\n\t Bank : " +
		a.Bank.String() +
		" \n}"
	return s
}
