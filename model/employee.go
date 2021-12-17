package model

type Student struct {
	ID          uint   `gorm:"primary_key"`
	Name        string `json:"name"`
	Dateofbirth string `json:"dateofbirth"`
	City        string `json:"city"`
	Designation string `json:"designation"`
	Joiningdate string `json:"joiningdate"`
}
