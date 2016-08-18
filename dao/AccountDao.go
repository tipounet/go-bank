package dao

import (
	"github.com/tipounet/go-bank/model"
	pg "gopkg.in/pg.v4"
)

// AccountDao : la dao d'un compte bancaire.
type AccountDao struct {
	DB *pg.DB
}

// GetAll : on récupère Tout
// FIXME : voir comment on fait pour les relations entre table. pour le moment epic fail :'
func (dao AccountDao) GetAll() (accounts []model.Account, err error) {
	err = dao.DB.Model(&accounts).
		Select()
	setAccountForeignData(accounts, dao.DB)
	return
}

// SearchByID : get account by ID
func (dao AccountDao) SearchByID(id int64) (account model.Account, err error) {
	err = dao.DB.Model(&account).
		Where("bankaccountid = ?", id).
		Select()
	u, _ := getUserByID(account.UserID, dao.DB)
	account.User = u
	b, _ := getBankByID(account.Bankid, dao.DB)
	account.Bank = b
	return
}

// FIXME : voir pour une recherche plus souple (like soundex etc ?)

// SearchByNumber search bank account from partial account number
func (dao AccountDao) SearchByNumber(accountNumber string) (accounts []model.Account, err error) {
	err = dao.DB.Model(&accounts).
		Where("accountnumber = ?", accountNumber).
		Select()
	setAccountForeignData(accounts, dao.DB)
	return
}

// SearchByUser search bank account for a user
func (dao AccountDao) SearchByUser(id int64) (accounts []model.Account, err error) {
	err = dao.DB.Model(&accounts).
		Where("userid = ?", id).
		Select()
	setAccountForeignData(accounts, dao.DB)
	return
}

// SearchByBank search bank account for a bank
func (dao AccountDao) SearchByBank(id int64) (accounts []model.Account, err error) {
	err = dao.DB.Model(&accounts).
		Where("bankid = ?", id).
		Select()
	setAccountForeignData(accounts, dao.DB)
	return
}

// Create : création d'un compte
func (dao AccountDao) Create(account *model.Account) error {
	return dao.DB.Create(account)
}

// Update : osef
func (dao AccountDao) Update(account *model.Account) error {
	return dao.DB.Update(account)
}

//Delete : suppression d'un compte
func (dao AccountDao) Delete(account *model.Account) error {
	return dao.DB.Delete(account)
}

// fonction privées
func getUserByID(id int64, db *pg.DB) (model.User, error) {
	udao := UserDao{
		DB: db,
	}
	return udao.GetByID(id)
}

// getBankByID : recherche d'une bank depuis son id
func getBankByID(id int64, db *pg.DB) (model.Bank, error) {
	bdao := BankDao{
		DB: db,
	}
	return bdao.GetByID(id)
}

// setUser : récuèpre un utilisateur depuis son ID
func setAccountForeignData(accounts []model.Account, db *pg.DB) {
	for i, a := range accounts {
		tmp := &accounts[i]
		u, e := getUserByID(a.UserID, db)
		if e == nil {
			tmp.User = u
		}

		b, eb := getBankByID(a.Bankid, db)
		if eb == nil {
			tmp.Bank = b
		}
	}
}
