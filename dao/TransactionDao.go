package dao

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/tipounet/go-bank/model"
)

// TransactionDao : la dao de la transaction !
type TransactionDao struct {
	DB *gorm.DB
}

// Get : recupère la liste des transaction sans distinction
func (dao TransactionDao) Read() (transactions []model.Transaction, err error) {
	err = dao.DB.
		Preload("Account.User").
		Preload("Account.Bank").
		Preload("TransactionType").
		Preload("Account").
		Order("posteddate desc, transactionid desc").Find(&transactions).Error
	return
}

// SearchByAccount : fonction de recherche par Account
func (dao TransactionDao) SearchByAccount(accountID int64) (t []model.Transaction, err error) {
	err = dao.DB.
		Preload("Account.User").
		Preload("Account.Bank").
		Preload("TransactionType").
		Preload("Account").
		Preload("TransactionType").
		Where("bankaccountid = ?", accountID).
		Order("posteddate desc").
		Find(&t).Error
	return
}

// SearchByDate : fonction de recherche par Posteddate ou userdate
func (dao TransactionDao) SearchByDate(date time.Time) (t []model.Transaction, err error) {
	err = dao.DB.
		Preload("Account.User").
		Preload("Account.Bank").
		Preload("TransactionType").
		Preload("Account").
		// FIXME : format date to force unckeck hours to_char ?
		Where("posteddate = ? or userdate = ?", date, date).
		Order("posteddate desc").
		Find(&t).Error
	return
}

// SearchByType : fonction de recherche par Type
func (dao TransactionDao) SearchByType(typeID int64) (t []model.Transaction, err error) {
	err = dao.DB.
		Preload("Account.User").
		Preload("Account.Bank").
		Preload("TransactionType").
		Preload("Account").
		Where("transaction_type_id = ?", typeID).
		Order("posteddate desc").
		Find(&t).Error
	return
}

// GetByID : retourne une transaction a partir de son id
func (dao TransactionDao) GetByID(id int64) (t model.Transaction, err error) {
	err = dao.DB.
		Preload("Account.User").
		Preload("Account.Bank").
		Preload("TransactionType").
		Preload("Account").
		First(&t, id).Error
	return
}

// Create : insertion dans la base
func (dao TransactionDao) Create(t *model.Transaction) (err error) {
	log.Printf("DAO :: insertion de %s\n", t)
	err = dao.DB.Set("gorm:save_associations", false).Create(t).Error
	return
}

// Update : mise à jour d'une transaction
// TODO : Voir pour prendre l'ID dans l'url vue que l'on fait un put (transaction/id)
func (dao TransactionDao) Update(t *model.Transaction, idOriginal int64) (err error) {
	// FIXME : le where double la condition car gorm prend quand même la pk l'objet du save :/
	// même problème avec Update
	err = dao.DB.Save(&t).Error
	//TODO : comment faire un update avec l'orm sans devoir se la faire à la main et en spécifiant moi même le prédicat ?
	return
}

//Delete : suppression Transaction
func (dao TransactionDao) Delete(t *model.Transaction) (err error) {
	err = dao.DB.Delete(t).Error
	return
}

// fonctions privées
// getAccountByID : retourne un compte depuis son id :-)~
func getAccountByID(id int64, db *gorm.DB) (model.Account, error) {
	adao := AccountDao{
		DB: db,
	}
	return adao.SearchByID(id)
}

// getTypeByID : recherche d'un type de transaction
func getTypeByID(id int64, db *gorm.DB) (model.TransactionType, error) {
	tdao := TransactionTypeDao{
		DB: db,
	}
	return tdao.GetByID(id)
}
