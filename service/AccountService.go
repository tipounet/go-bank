package service

import (
	"github.com/tipounet/go-bank/dao"
	"github.com/tipounet/go-bank/model"
)

// TODO : Search by user / bank modify to searchBy userName / bankName (with like ;) )

// AccountService service métier de gestion des comte bancaire
type AccountService struct {
	Dao *dao.AccountDao
}

// SearchByID : search account by account id
func (service AccountService) SearchByID(id int64) (account model.Account, err error) {
	return service.Dao.SearchByID(id)
}

// SearchByNumber : search account by account number
func (service AccountService) SearchByNumber(accountNumber string) (account []model.Account, err error) {
	return service.Dao.SearchByNumber(accountNumber)
}

// SearchByUser : search account by user id
func (service AccountService) SearchByUser(id int64) (account []model.Account, err error) {
	return service.Dao.SearchByUser(id)
}

// SearchByBank : search account by bankid
func (service AccountService) SearchByBank(id int64) (account []model.Account, err error) {
	return service.Dao.SearchByBank(id)
}

// Create : création d'un nouveau compte
func (service AccountService) Create(account *model.Account) (err error) {
	return service.Dao.Create(account)
}

// Read : liste de tout les comtes
func (service AccountService) Read() (accounts []model.Account, err error) {
	return service.Dao.GetAll()
}

// Update : mise à jour d'un compte
func (service AccountService) Update(account *model.Account) (err error) {
	return service.Dao.Update(account)
}

// Delete : suppression d'un compte
func (service AccountService) Delete(account *model.Account) (err error) {
	return service.Dao.Delete(account)
}
