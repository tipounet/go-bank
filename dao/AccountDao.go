package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/tipounet/go-bank/model"
)

// AccountDao : la dao d'un compte bancaire.
type AccountDao struct {
	DB *gorm.DB
}

// GetAll : on récupère Tout
// FIXME : voir comment on fait pour les relations entre table. pour le moment epic fail :'
func (dao AccountDao) GetAll() (accounts []model.Account, err error) {
	err = dao.DB.Preload("User").Preload("Bank").Order("accountnumber asc").Find(&accounts).Error
	// setAccountForeignData(accounts, dao.DB)
	return
}

// SearchByID : get account by ID
func (dao AccountDao) SearchByID(id int64) (account model.Account, err error) {
	err = dao.DB.Preload("User").Preload("Bank").Order("accountnumber asc").First(&account, id).Error
	return
}

// FIXME : voir pour une recherche plus souple (like soundex etc ?)

// SearchByNumber search bank account from partial account number
func (dao AccountDao) SearchByNumber(accountNumber string) (accounts []model.Account, err error) {
	err = dao.DB.Preload("User").Preload("Bank").Where("accountnumber = ?", accountNumber).Order("accountnumber asc").Find(&accounts).Error
	return
}

// SearchByUser search bank account for a user
// FIXME : rename la focntion  serachByUserID
func (dao AccountDao) SearchByUser(id int64) (accounts []model.Account, err error) {
	err = dao.DB.Preload("User").Preload("Bank").Where("userid = ?", id).Find(&accounts).Error
	return
}

// SearchByBank search bank account for a bank
func (dao AccountDao) SearchByBank(id int64) (accounts []model.Account, err error) {
	err = dao.DB.Preload("User").Preload("Bank").Where("bankid = ?", id).Find(&accounts).Error
	return
}

// Create : création d'un compte
// FIXME : voir si je dois ounon garder le set avec la suppression de la sauvegarde des associations
func (dao AccountDao) Create(account *model.Account) (err error) {
	err = dao.DB.Set("gorm:save_associations", false).Create(account).Error
	return
}

// Update : osef
func (dao AccountDao) Update(account *model.Account) (err error) {
	err = dao.DB.Set("gorm:save_associations", false).Save(account).Error
	return
}

//Delete : suppression d'un compte
func (dao AccountDao) Delete(account *model.Account) (err error) {
	err = dao.DB.Delete(account).Error
	return
}

// fonction privées
func getUserByID(id int64, db *gorm.DB) (model.User, error) {
	udao := UserDao{
		DB: db,
	}
	return udao.GetByID(id)
}

// getBankByID : recherche d'une bank depuis son id
func getBankByID(id int64, db *gorm.DB) (model.Bank, error) {
	bdao := BankDao{
		DB: db,
	}
	return bdao.GetByID(id)
}
