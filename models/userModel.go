package models

import (
	"errors"
	"strings"
	"wishlist/entities"
	"wishlist/helper"
	// "wishlist/middleware"

	// "github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type WhistUser entities.WhistUser

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *WhistUser) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *WhistUser) Validate(action string) error {
	switch strings.ToLower(action) {
	case "create":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Role == "" {
			return errors.New("Required Role")
		}
		if u.EmailAddress == "" {
			return errors.New("Required Email")
		}
		if u.PhoneNumber == "" {
			return errors.New("Required Phone Number")
		}
		return nil

	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.EmailAddress == "" {
			return errors.New("Required Email")
		}
		return nil

	default:
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.EmailAddress == "" {
			return errors.New("Required Email")
		}
		return nil
	}
}

func (u *WhistUser) SaveUsers() (*WhistUser, error) {
	var err error
	u.Uid = helper.GENERATEUUID()
	u.Status = true

	err = db.Debug().Create(&u).Error
	if err != nil {
		return &WhistUser{}, err
	}
	return u, nil
}

// func (u *WhistUser) FindAllUsers(db *gorm.DB, middleware *middleware.Access) (*[]WhistUser, error) {
// 	var err error
// 	var users []WhistUser

// 	err = db.Debug().Where("status != ?", false).Find(&users).Error
// 	if err != nil {
// 		return &[]WhistUser{}, err
// 	}
// 	return &users, err
// }

// func (u *WhistUser) FindUsersByID(db *gorm.DB, uid string, middleware *middleware.Access) (*[]WhistUser, error) {
// 	var err error
// 	users := []WhistUser{}

// 	err = db.Debug().Model(&WhistUser{}).Where("uid = ? and status != ?", uid, false).Find(&users).Error
// 	if err != nil {
// 		return &[]WhistUser{}, err
// 	}
// 	return &users, nil
// }

// func (u *WhistUser) SaveUpdateUsers(db *gorm.DB, uid string, middleware *middleware.Access) (*WhistUser, error) {
// 	err := db.Debug().Model(&WhistUser{}).Where("uid = ?", uid).Update(&u).Error
// 	if err != nil {
// 		return &WhistUser{}, err
// 	}
// 	return u, nil
// }

// func (u *WhistUser) SaveDeleteUsers(db *gorm.DB, uid string, middleware *middleware.Access) (*WhistUser, error) {
// 	u.Status = false
// 	err := db.Debug().Model(&WhistUser{}).Where("uid = ?", uid).Update(&u).Error
// 	if err != nil {
// 		return &WhistUser{}, err
// 	}
// 	return u, nil
// }
