package model

import (
	"strconv"
	"time"
)

// Transaction c'est la table contient les recettes et dépenses
type Transaction struct {
	Transactionid     int64     `gorm:"primary_key" json:"id"`
	Description       string    `json:"description"`
	Posteddate        time.Time `json:"Posteddate"`
	Userdate          time.Time `json:"userdate"`
	Fiid              string    `json:"fiid"`
	Amount            float64   `json:"amount"`
	BankaccountID     int64     `gorm:"column:bankaccountid" json:"accountID"`
	Account           Account   `gorm:"ForeignKey:Bankaccountid;AssociationForeignKey:BankaccountID" json:"account" `
	TransactionTypeID int64     `json:"typeID" gorm:"column:transaction_type_id"`
	// gorm:"ForeignKey:le_nom_de_la_colonne_dans_la_table;AssociationForeignKey:leNomDansLobjet
	TransactionType TransactionType `gorm:"ForeignKey:transaction_type_id;AssociationForeignKey:TransactionTypeID" json:"type"`
}

// TableName : permet d'indiquer le nom de la table sinon gorm utilise le pluriel (users)
func (Transaction) TableName() string {
	return "transaction"
}

func (t Transaction) String() string {
	retour := "Transaction {\n"
	retour = retour + "\tTransactionid : " + strconv.FormatInt(t.Transactionid, 10) + "\n"
	retour = retour + "\tDescription : " + t.Description + "\n"
	// retour = retour + "\tPosteddate : " + t.Posteddate + "\n"
	// retour = retour + "\tUserdate : " + strconv.FormatTime(t.Userdate) + "\n"
	retour = retour + "\tFiid : " + t.Fiid + "\n"
	// retour = retour + "\tAmount : " + strconv.FormatFloat(t.Amount, 10) + "\n"
	retour = retour + "\tBankaccountid : " + strconv.FormatInt(t.BankaccountID, 10) + "\n"
	retour = retour + "\tAccount : " + t.Account.String() + "\n"
	retour = retour + "\tTransactionTypeID : " + strconv.FormatInt(t.TransactionTypeID, 10) + "\n"
	retour = retour + "}"
	return retour
}
