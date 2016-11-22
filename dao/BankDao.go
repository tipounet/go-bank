package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/tipounet/go-bank/model"
)

// BankDao : la dao de la bank !
type BankDao struct {
	DB *gorm.DB
}

// Get ça c'est le get de BankDao
func (dao BankDao) Get() (banks []model.Bank, err error) {
	err = dao.DB.Order("name asc").Find(&banks).Error
	return
}

// GetByID : return a Bank from id
func (dao BankDao) GetByID(id int64) (bank model.Bank, err error) {
	// FIXME : pas d'erreur quand trouve rien et en plus pas de soucis on retourne bank qui est valide même avec rien :-()~
	err = dao.DB.First(&bank, id).Error
	return
}

// SearchByName : retourne les banques qui correspodnent au nom name
// FIXME : voir pour un recherche plus poussée dans le nom (genre que des lettres comme dans idea)
func (dao BankDao) SearchByName(name string) (banks []model.Bank, err error) {
	err = dao.DB.Where("lower(name) like concat('%',lower(?),'%')", name).
		Order("name desc").
		Find(&banks).Error
	return
}

// Create : jesus ?
func (dao BankDao) Create(bank *model.Bank) (err error) {
	err = dao.DB.Set("gorm:save_associations", false).Create(bank).Error
	return
}

// Update : osef
func (dao BankDao) Update(bank *model.Bank) (err error) {
	err = dao.DB.Save(bank).Error
	return
}

//Delete : suppression d'une bank
func (dao BankDao) Delete(bank *model.Bank) (err error) {
	err = dao.DB.Delete(bank).Error
	return
}
