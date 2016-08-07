package model

// TransactionType : représente un type de transaction
type TransactionType struct {
	TableName struct{} `sql:"transaction_type" json:"-"`
	ID        int64    `sql:"transaction_type_id,pk" json:"id"`
	Name      string   `sql:"transaction_name" json:"name"`
}
