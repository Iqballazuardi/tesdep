package usecase_test

import (
	"deploy/features/user"
	"deploy/features/user/mocks"
	"deploy/features/user/usecase"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	// fixture
	repo := mocks.NewRepository(t)
	ul := usecase.New(repo)
	successCaseData := user.Core{Nama: "jerry", HP: "12345", Password: "tonohaha577"}

	t.Run("Sukses login", func(t *testing.T) {
		repo.On("Login", successCaseData.HP, successCaseData.Password).Return(user.Core{Nama: "jerry", HP: "12345"}, nil).Once()
		result, err := ul.Login("12345", "tonohaha577")

		assert.Nil(t, err)
		assert.Equal(t, "12345", result.HP)
		assert.Equal(t, "jerry", result.Nama)
		repo.AssertExpectations(t)
	})

	t.Run("Password salah", func(t *testing.T) {
		repo.On("Login", successCaseData.HP, "tonohaha").Return(user.Core{}, errors.New("password tidak sesuai")).Once()
		result, err := ul.Login("12345", "tonohaha")

		assert.Error(t, err)
		assert.ErrorContains(t, err, "sesuai")
		assert.Empty(t, result.Nama)
		repo.AssertExpectations(t)
	})

	t.Run("Data tidak ditemukan", func(t *testing.T) {
		repo.On("Login", "6789", "tonohaha").Return(user.Core{}, errors.New("data tidak ditemukan")).Once()
		result, err := ul.Login("6789", "tonohaha")

		assert.Error(t, err)
		assert.ErrorContains(t, err, "data tidak ditemukan")
		assert.Empty(t, result.Nama)
		repo.AssertExpectations(t)
	})

	t.Run("Kesalahan pada server", func(t *testing.T) {
		repo.On("Login", successCaseData.HP, "tonohaha").Return(user.Core{}, errors.New("column not exist")).Once()
		result, err := ul.Login("12345", "tonohaha")

		assert.Error(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Empty(t, "", result.Nama)
		repo.AssertExpectations(t)
	})
}

func TestRegister(t *testing.T) {
	repo := mocks.NewRepository(t)
	ul := usecase.New(repo)
	insertData := user.Core{
		Nama:     "jerry",
		HP:       "12345",
		Password: "alta123",
	}

	t.Run("Sukses register", func(t *testing.T) {
		repo.On("Insert", insertData).Return(insertData, nil).Once()
		err := ul.Register(insertData)

		assert.Nil(t, err)
	})

	t.Run("Gagal register", func(t *testing.T) {
		repo.On("Insert", insertData).Return(user.Core{}, errors.New("too many values")).Once()
		err := ul.Register(insertData)

		assert.Error(t, err)
	})
}
