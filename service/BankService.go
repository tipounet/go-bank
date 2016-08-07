package service

import (
	"github.com/tipounet/go-bank/dao"
	"github.com/tipounet/go-bank/model"
)

// BankService service métier de gestion des bank
type BankService struct {
	Dao *dao.BankDao
}

// Get : toutes les bank du monde
func (service BankService) Get() (banks []model.Bank, err error) {
	return service.Dao.Get()
}

// Search Recherche d'utilsiateur
func (service BankService) Search(id int64) (bank model.Bank, err error) {
	return service.Dao.GetByID(id)
}

// SearchPartialName Recherche d'utilsiateur
func (service BankService) SearchPartialName(name string) (bank []model.Bank, err error) {
	return service.Dao.SearchByName(name)
}

// Create : création d'une nouvelle banque
func (service BankService) Create(bank *model.Bank) (err error) {
	return service.Dao.Create(bank)
}

// Read : liste de tout les banques
func (service BankService) Read() (banks []model.Bank, err error) {
	return service.Dao.Get()
}

// Update : mise à jour d'une banques
func (service BankService) Update(bank *model.Bank) (err error) {
	return service.Dao.Update(bank)
}

// Delete : suppression d'une banques
func (service BankService) Delete(bank *model.Bank) (err error) {
	return service.Dao.Delete(bank)
}
