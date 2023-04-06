package handler

type RegisterInput struct {
	Nama     string `json:"nama"`
	HP       string `json:"hp"`
	Password string `json:"password"`
}

type LoginInput struct {
	HP       string `json:"hp"`
	Password string `json:"password"`
}
