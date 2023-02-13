package db_manager

import (
	"work/model"

	"gorm.io/gorm"
)

type DbManager struct {
	db *gorm.DB
}

func NewDbManager(database *gorm.DB) *DbManager {
	return &DbManager{db: database}
}

// func (d *DbManager) First() {
// 	fmt.Println("conectet")
// }


//save order to database
func (s *DbManager) AddOrder(order *model.Order) error {
	return s.db.Create(order).Error
}
