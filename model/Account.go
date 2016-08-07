package model

// Account : c'est un compte bancaire a quelqu'un
type Account struct {
	TableName     struct{} `sql:"bankaccount" json:"-"` //json ignore
	Bankaccountid int64    `sql:",pk" json:"id"`
	Accountnumber string   `json:"number"`
	// Has one does not work
	UserID int64 `sql:"userid" json:"-,omitempty"`
	User   User  `pg:",joinFK:userid" json:"user"` // des chose avec des {` pour donner la clef de jointure ?`}
	Bankid int64 `json:"-,omitempty"`
	Bank   Bank  `pg:",joinFK:bankid" json:"bank"`
}
