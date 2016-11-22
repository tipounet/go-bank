package model

// TransactionType : représente un type de transaction
type TransactionType struct {
	TransactionTypeID int64  `gorm:"primary_key;column:transaction_type_id" json:"id"`
	Name              string `gorm:"column:transaction_name" json:"name"`
}

// TableName : permet d'indiquer le nom de la table sinon gorm utilise le pluriel (users)
func (TransactionType) TableName() string {
	return "transaction_type"
}
