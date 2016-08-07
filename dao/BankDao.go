package dao

import (
	pg "gopkg.in/pg.v4"
	"github.com/tipounet/go-bank/model"
)

// BankDao : la dao de la bank !
type BankDao struct {
	DB *pg.DB
}

// Get ça c'est le get de BankDao
func (dao BankDao) Get() (banks []model.Bank, err error) {
	err = dao.DB.Model(&banks).
		Order("name asc").
		Select()
	return banks, err
}

// GetByID : return a Bank from id
func (dao BankDao) GetByID(id int64) (bank model.Bank, err error) {
	err = dao.DB.Model(&bank).Where("bankid=?", id).Select()
	return
}

// SearchByName : retourne les banques qui correspodnent au nom name
// FIXME : voir pour un recherche plus poussée dans le nom (genre que des lettres comme dans idea)
func (dao BankDao) SearchByName(name string) (banks []model.Bank, err error) {
	err = dao.DB.Model(&banks).
		Where("lower(name) like concat('%',lower(?),'%')", name).
		Order("name desc").
		Select()
	return
}

// Create : jesus ?
func (dao BankDao) Create(bank *model.Bank) error {
	return dao.DB.Create(bank)
}

// Update : osef
func (dao BankDao) Update(bank *model.Bank) error {
	return dao.DB.Update(bank)
}

//Delete : suppression d'une bank
func (dao BankDao) Delete(bank *model.Bank) error {
	return dao.DB.Delete(bank)
}
