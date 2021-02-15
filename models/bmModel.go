package models

import (
	// "github.com/jinzhu/gorm"
	"wishlist/entities"
	// "wishlist/middleware"
)

type WhistBm entities.WhistBm

// func (u *WhistBm) SaveBm(db *gorm.DB) (*WhistBm, error) {
// 	u.Status = "entered"
// 	err := db.Debug().Create(&u).Error
// 	if err != nil {
// 		return &WhistBm{}, err
// 	}
// 	return u, nil
// }

// func (u *WhistBm) FindAllBm(db *gorm.DB, middleware *middleware.Access) (*[]WhistBm, error) {
// 	var data []WhistBm
// 	err := db.Debug().Where("is_deleted != ?", true).Find(&data).Error
// 	if err != nil {
// 		return &[]WhistBm{}, err
// 	}
// 	return &data, err
// }

// func (u *WhistBm) FindBmByID(db *gorm.DB, ID string, middleware *middleware.Access) (*[]WhistBm, error) {
// 	data := []WhistBm{}
// 	err := db.Debug().Model(&WhistBm{}).Where("id = ? and is_deleted != ?", ID, true).Find(&data).Error
// 	if err != nil {
// 		return &[]WhistBm{}, err
// 	}
// 	return &data, nil
// }

// func (u *WhistBm) SaveUpdateBm(db *gorm.DB, ID string, middleware *middleware.Access) (*WhistBm, error) {
// 	err := db.Debug().Model(&WhistBm{}).Where("id = ?", ID).Update(&u).Error
// 	if err != nil {
// 		return &WhistBm{}, err
// 	}
// 	return u, nil
// }
