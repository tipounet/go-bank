package model

// Bank : une banque !
type Bank struct {
	TableName struct{} `sql:"bank" json:"-"`
	Bankid    int64    `sql:",pk" json:"id"`
	Name      string   `json:"name"`
}
