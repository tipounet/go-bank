package dao

// FIXME : Mapper toutes les opérations d'un crud (manque le search en fait, avec un prédicat dynamique ?)
import (
	pg "gopkg.in/pg.v4"
	"github.com/tipounet/go-bank/model"
)

// TransactionTypeDao : la dao du type de transation
type TransactionTypeDao struct {
	DB *pg.DB
}

// Get : retourne tout les types de transaction
func (dao TransactionTypeDao) Get() (retour []model.TransactionType, err error) {
	err = dao.DB.Model(&retour).
		Order("transaction_name asc").
		Select()
	return
}

// GetByID : retourne un type de transaction en fonction de son ID
func (dao TransactionTypeDao) GetByID(id int64) (retour model.TransactionType, err error) {
	err = dao.DB.Model(&retour).
		Where("transaction_type_id=?", id).
		Select()
	return
}

// Create :Création d'un type de transaction (ça va pas arriver tout les jour ;)
func (dao TransactionTypeDao) Create(tt *model.TransactionType) error {
	return dao.DB.Create(tt)
}

// Update : Mise à jour
func (dao TransactionTypeDao) Update(tt *model.TransactionType) error {
	return dao.DB.Update(tt)
}

//Delete : suppression d'un type de transaction
func (dao TransactionTypeDao) Delete(tt *model.TransactionType) error {
	return dao.DB.Delete(tt)
}
