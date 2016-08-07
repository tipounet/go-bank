package service

import (
	"github.com/tipounet/go-bank/dao"
	"github.com/tipounet/go-bank/model"
)

// TransactionTypeService service métier de gestion des types de transactions
type TransactionTypeService struct {
	Dao *dao.TransactionTypeDao
}

// SearchByID Recherche d'un type de transaction
func (service TransactionTypeService) SearchByID(id int64) (transactiontype model.TransactionType, err error) {
	return service.Dao.GetByID(id)
}

// Create : création d'un nouveau type de transaction
func (service TransactionTypeService) Create(transactiontype *model.TransactionType) (err error) {
	return service.Dao.Create(transactiontype)
}

// Read : liste de tout les type de transaction
func (service TransactionTypeService) Read() (transactiontypes []model.TransactionType, err error) {
	return service.Dao.Get()
}

// Update : mise à jour d'un type de transaction
func (service TransactionTypeService) Update(transactiontype *model.TransactionType) (err error) {
	return service.Dao.Update(transactiontype)
}

// Delete : suppression d'un type de transaction
func (service TransactionTypeService) Delete(transactiontype *model.TransactionType) (err error) {
	return service.Dao.Delete(transactiontype)
}
