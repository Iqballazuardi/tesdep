package user

import "github.com/labstack/echo/v4"

type Core struct {
	Nama     string
	HP       string
	Password string
}

type Handler interface {
	Login() echo.HandlerFunc
	Register() echo.HandlerFunc
}

type UseCase interface {
	Login(hp string, password string) (Core, error)
	Register(newUser Core) error
}

type Repository interface {
	Insert(newUser Core) (Core, error)
	Login(hp string, password string) (Core, error)
}
