package handler

type LoginResponse struct {
	HP    string `json:"hp"`
	Nama  string `json:"nama"`
	Token string `json:"token"`
}
