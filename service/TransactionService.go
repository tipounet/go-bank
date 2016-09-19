package service

import (
	"time"

	"github.com/tipounet/go-bank/dao"
	"github.com/tipounet/go-bank/model"
)

// TransactionService service métier de gestion des transactions
type TransactionService struct {
	Dao *dao.TransactionDao
}

// SearchByID Recherche d'une transaction à partir de son ID
func (service TransactionService) SearchByID(id int64) (model.Transaction, error) {
	return service.Dao.GetByID(id)
}

// SearchByAccount : Recherche de transaction par Account
func (service TransactionService) SearchByAccount(id int64) ([]model.Transaction, error) {
	return service.Dao.SearchByAccount(id)
}

// SearchByDate : Recherche de transaction par Date
func (service TransactionService) SearchByDate(date time.Time) ([]model.Transaction, error) {
	return service.Dao.SearchByDate(date)
}

// SearchByType : Recherche de transaction par Type
func (service TransactionService) SearchByType(id int64) ([]model.Transaction, error) {
	return service.Dao.SearchByType(id)
}

// Create : création d'une nouvelle transaction
func (service TransactionService) Create(transaction *model.Transaction) error {
	return service.Dao.Create(transaction)
}

// Read : liste de toutes les transactions
func (service TransactionService) Read() ([]model.Transaction, error) {
	return service.Dao.Read()
}

// Update : mise à jour d'une transaction
func (service TransactionService) Update(transaction *model.Transaction, idOriginal int64) error {
	return service.Dao.Update(transaction, idOriginal)
}

// Delete : suppression d'une transaction
func (service TransactionService) Delete(transaction *model.Transaction) error {
	return service.Dao.Delete(transaction)
}
