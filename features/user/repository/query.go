package repository

import (
	"errors"

	"deploy/features/user"
	"deploy/helper"

	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type userModel struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.Repository {
	return &userModel{
		db: db,
	}
}

func (um *userModel) Insert(newUser user.Core) (user.Core, error) {
	var inputUser = User{}
	hashedPassword, err := helper.GenerateHashedPassword(newUser.Password)
	if err != nil {
		log.Error("terjadi error saat melakukan hashing", err.Error())
		return user.Core{}, err
	}
	inputUser.HP = newUser.HP
	inputUser.Nama = newUser.Nama
	inputUser.Password = hashedPassword
	inputUser.ID = uuid.New().String()

	if err := um.db.Table("users").Create(&inputUser).Error; err != nil {
		log.Error("Terjadi error saat create user", err.Error())
		return user.Core{}, err
	}

	return newUser, nil
}

func (um *userModel) Login(hp, password string) (user.Core, error) {
	res := User{}
	if err := um.db.Where("hp = ? ", hp).Find(&res).Error; err != nil {
		log.Error("Terjadi error saat login user", err.Error())
		return user.Core{}, err
	}

	if res.HP == "" {
		log.Error("Data tidak ditemukan")
		return user.Core{}, errors.New("data tidak ditemukan")
	}

	if !helper.ComparePassword(res.Password, password) {
		log.Error("Password salah")
		return user.Core{}, errors.New("password tidak sesuai")
	}

	return user.Core{Nama: res.Nama, HP: res.HP}, nil
}

func (um *userModel) GetAllUser() ([]User, error) {
	res := []User{}

	if err := um.db.Select("hp, nama, id").Find(&res).Error; err != nil {
		log.Error("Terjadi error saat select user ", err.Error())
		return nil, err
	}

	return res, nil
}
