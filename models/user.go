package models

type User struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	NamaLengkap string `gorm:"varchar(255)" json:"nama_lengkap"`
	UserName    string `gorm:"varchar(255)" json:"user_name"`
	Password    string `gorm:"varchar(255)" json:"password"`
}
