package model

import (
	"time"
)

// Transaction c'est la table contient les recettes et dépenses
type Transaction struct {
	TableName         struct{}        `sql:"transaction" json:"-"`
	Transactionid     int64           `sql:",pk" json:"id"`
	TransactionTypeID int64           `sql:"transaction_type_id" json:"-"`
	Description       string          `json:"description"`
	Posteddate        time.Time       `json:"Posteddate"`
	Userdate          time.Time       `json:"userdate"`
	Fiid              string          `json:"fiid"`
	Amount            float64         `json:"amount"`
	Bankaccountid     string          `json:"-"`
	Account           Account         `json:"account"`
	Type              TransactionType `json:"type"`
}
