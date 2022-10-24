package model

import "time"

type BaseModel struct {
	Id        string `json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Nip      string `json:"nip"`
	Address  string `json:"address"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

var Users = []User{}
