package handler

import (
	"deploy/features/user"
	"deploy/helper"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type userController struct {
	service user.UseCase
}

func New(us user.UseCase) user.Handler {
	return &userController{
		service: us,
	}
}

func (uc *userController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := RegisterInput{}
		if err := c.Bind(&input); err != nil {
			c.Logger().Error("terjadi kesalahan bind", err.Error())
			return c.JSON(helper.ReponsFormat(http.StatusBadRequest, "terdapat kesalahan input dari user", nil))
		}

		err := uc.service.Register(user.Core{HP: input.HP, Nama: input.Nama, Password: input.Password})

		if err != nil {
			return c.JSON(helper.ReponsFormat(http.StatusInternalServerError, err.Error(), nil))
		}

		return c.JSON(helper.ReponsFormat(http.StatusCreated, "sukses menambahkan data", nil))
	}
}

func (uc *userController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginInput
		if err := c.Bind(&input); err != nil {
			c.Logger().Error("terjadi kesalahan bind", err.Error())
			return c.JSON(helper.ReponsFormat(http.StatusBadRequest, "terdapat kesalahan input dari user", nil))
		}

		res, err := uc.service.Login(input.HP, input.Password)
		if err != nil {
			code := http.StatusInternalServerError
			if strings.Contains(err.Error(), "sesuai") || strings.Contains(err.Error(), "ditemukan") {
				code = http.StatusBadRequest
			}
			return c.JSON(helper.ReponsFormat(code, err.Error(), nil))
		}

		var result = new(LoginResponse)
		token := helper.GenerateJWT(res.HP)
		result.Nama = res.Nama
		result.HP = res.HP
		result.Token = token

		return c.JSON(helper.ReponsFormat(http.StatusOK, "sukses login, gunakan token ini pada akses api selanjutnya : ", result))
	}
}
