package models

import "gorm.io/gorm"

type Rate struct {
	gorm.Model
	Days     string `gorm:"size:125;not null;unique"`
	Times    string `gorm:"size:9;not null;unique"`
	TimeZone string `gorm:"size:125;not null;unique"`
	Price    int
} //end struct
