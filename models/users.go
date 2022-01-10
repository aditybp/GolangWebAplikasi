package models

type User struct {
	Id       uint   `json:"id_user"`
	Nama     string `json:"nama_user" `
	Email    string `json:"email_user"`
	Password string `json:"password_user"`
}
