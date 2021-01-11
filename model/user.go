package model

import "gorm.io/gorm"

// User for Model Table Database
type User struct {
	gorm.Model
	Username    string
	Password    string
	NamaLengkap string
	Foto        string
}
