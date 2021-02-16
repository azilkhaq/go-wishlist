package models

import (
	"wishlist/entities"
)

type WhistBm entities.WhistBm

func (u *WhistBm) SaveBm() (*WhistBm, error) {
	
	u.Status = "entered"
	err := db.Debug().Create(&u).Error
	if err != nil {
		return &WhistBm{}, err
	}
	return u, nil
}

func SaveAllBm() (*[]WhistBm, error) {

	var data []WhistBm
	err := db.Debug().Where("is_deleted != ?", true).Find(&data).Error
	if err != nil {
		return &[]WhistBm{}, err
	}
	return &data, err
}

func SaveSingleBm(ID string) (*[]WhistBm, error) {

	data := []WhistBm{}
	err := db.Debug().Model(&WhistBm{}).Where("id = ? and is_deleted != ?", ID, true).Find(&data).Error
	if err != nil {
		return &[]WhistBm{}, err
	}
	return &data, nil
}

func (u *WhistBm) SaveUpdateBm(ID string) (*WhistBm, error) {

	err := db.Debug().Model(&WhistBm{}).Where("id = ?", ID).Update(&u).Error
	if err != nil {
		return &WhistBm{}, err
	}
	return u, nil
}
