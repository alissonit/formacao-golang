package models

import "gorm.io/gorm"

// https://gorm.io/docs/models.html
type Student struct {
	gorm.Model
	Name string `json:"name"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
}

var Students []Student
