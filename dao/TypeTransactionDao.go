package dao

// FIXME : Mapper toutes les opérations d'un crud (manque le search en fait, avec un prédicat dynamique ?)
import (
	"github.com/jinzhu/gorm"
	"github.com/tipounet/go-bank/model"
)

// TransactionTypeDao : la dao du type de transation
type TransactionTypeDao struct {
	DB *gorm.DB
}

// Get : retourne tout les types de transaction
func (dao TransactionTypeDao) Get() (retour []model.TransactionType, err error) {
	err = dao.DB.Order("transaction_name asc").Find(&retour).Error
	return
}

// GetByID : retourne un type de transaction en fonction de son ID
func (dao TransactionTypeDao) GetByID(id int64) (retour model.TransactionType, err error) {
	err = dao.DB.First(&retour, "transaction_type_id=?", id).Error
	return
}

// Create :Création d'un type de transaction (ça va pas arriver tout les jour ;)
func (dao TransactionTypeDao) Create(tt *model.TransactionType) (err error) {
	err = dao.DB.Set("gorm:save_associations", false).Create(tt).Error
	return
}

// Update : Mise à jour
func (dao TransactionTypeDao) Update(tt *model.TransactionType) (err error) {
	err = dao.DB.Save(tt).Error
	return
}

//Delete : suppression d'un type de transaction
func (dao TransactionTypeDao) Delete(tt *model.TransactionType) (err error) {
	err = dao.DB.Delete(tt).Error
	return
}
