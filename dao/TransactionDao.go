package dao

import (
	"time"

	pg "gopkg.in/pg.v4"

	"github.com/tipounet/go-bank/model"
)

// TransactionDao : la dao de la transaction !
type TransactionDao struct {
	DB *pg.DB
}

// Get : recupère la liste des transaction sans distinction
func (dao TransactionDao) Read() (transactions []model.Transaction, err error) {
	err = dao.DB.Model(&transactions).
		Order("posteddate desc").
		Select()

	setTransactionForeignData(transactions, dao.DB)
	return
}

// SearchByAccount : fonction de recherche par Account
func (dao TransactionDao) SearchByAccount(accountID int64) (t []model.Transaction, err error) {
	err = dao.DB.Model(&t).
		Where("bankaccountid = ?", accountID).
		Order("posteddate desc").
		Select()
	setTransactionForeignData(t, dao.DB)
	return
}

// SearchByDate : fonction de recherche par Posteddate ou userdate
func (dao TransactionDao) SearchByDate(date time.Time) (t []model.Transaction, err error) {
	err = dao.DB.Model(&t).
		// FIXME : format date to force unckeck hours
		Where("posteddate = ? or userdate = ?", date).
		Order("posteddate desc").
		Select()
	setTransactionForeignData(t, dao.DB)
	return
}

// SearchByType : fonction de recherche par Type
func (dao TransactionDao) SearchByType(typeID int64) (t []model.Transaction, err error) {
	err = dao.DB.Model(&t).
		Where("transaction_type_id = ?", typeID).
		Order("posteddate desc").
		Select()
	setTransactionForeignData(t, dao.DB)
	return
}

// GetByID : retourne une transaction a partir de son id
func (dao TransactionDao) GetByID(id int64) (t model.Transaction, err error) {
	err = dao.DB.Model(&t).Where("transactionid=?", id).Select()
	a, e := getAccountByID(t.Transactionid, db)
	if e == nil {
		t.Account = a
	}

	ty, et := getTypeByID(t.TransactionTypeID, db)
	if et == nil {
		t.Type = ty
	}
	return
}

// Create : insertion dans la base
func (dao TransactionDao) Create(t *model.Transaction) error {
	return dao.DB.Create(t)
}

// Update : osef
func (dao TransactionDao) Update(t *model.Transaction) error {
	return dao.DB.Update(&t)
}

//Delete : suppression Transaction
func (dao TransactionDao) Delete(t *model.Transaction) error {
	return dao.DB.Delete(t)
}

// fonctions privées
// getAccountByID : retourne un compte depuis son id :-)~
func getAccountByID(id int64, db *pg.DB) (model.Account, error) {
	adao := AccountDao{
		DB: db,
	}
	return adao.SearchByID(id)
}

// getTypeByID : recherche d'un type de transaction
func getTypeByID(id int64, db *pg.DB) (model.TransactionType, error) {
	tdao := TransactionTypeDao{
		DB: db,
	}
	return tdao.GetByID(id)
}

// setTransactionForeignData : ajout des compte bancaire correspondant aux transactions passées en paramètre. On travail direct sur la première liste histoire d'éviter l'emprunte mémoire trop importante.
func setTransactionForeignData(transactions []model.Transaction, db *pg.DB) {
	for i, t := range transactions {
		tmp := &transactions[i]
		a, e := getAccountByID(t.Transactionid, db)
		if e == nil {
			tmp.Account = a
		}

		ty, et := getTypeByID(t.Transactionid, db)
		if et == nil {
			tmp.Type = ty
		}
	}
}
