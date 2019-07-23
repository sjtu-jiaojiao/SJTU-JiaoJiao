package db

import (
	"time"
)

type Test struct {
	ID       int32  `gorm:"auto_increment;primary_key"`
	TestName string `gorm:"size:100"`
}

type User struct {
	ID          int32  `gorm:"auto_increment;primary_key"`
	UserName    string `gorm:"size:32;not null"`
	AvatarId    string `gorm:"type:char(24)"`
	Telephone   string `gorm:"type:char(11)"`
	StudentId   string `gorm:"size:32;index;not null"`
	StudentName string `gorm:"size:32;not null"`
	Status      int32  `gorm:"default:1;not null"`
	Role        int32  `gorm:"default:1;not null"`
}

type SellInfo struct {
	ID          int32     `gorm:"auto_increment;primary_key"`
	Status      int32     `gorm:"default:1;not null"`
	ReleaseTime time.Time `gorm:"not null"`
	ValidTime   time.Time `gorm:"not null"`
	UserId      int32     `gorm:"not null;index"`
	GoodId      int32     `gorm:"not null;index"`
}

type BuyInfo struct {
	ID          int32     `gorm:"auto_increment;primary_key"`
	Status      int32     `gorm:"default:1;not null"`
	ReleaseTime time.Time `gorm:"not null"`
	ValidTime   time.Time `gorm:"not null"`
	UserId      int32     `gorm:"not null;index"`
	GoodId      int32     `gorm:"not null;index"`
}

type Transaction struct {
	ID         int32     `gorm:"auto_increment;primary_key"`
	CreateTime time.Time `gorm:"not null"`
	InfoId     int32     `gorm:"not null"`
	UserId     int32     `gorm:"not null"`
	Status     int32     `gorm:"not null;default:1"`
}

type Good struct {
	ID          int32   `gorm:"auto_increment;primary_key"`
	GoodName    string  `gorm:"size:128;not null"`
	Price       float64 `gorm:"default:0;not null"`
	Description string  `gorm:"type:text"`
	ContentId   string  `gorm:"type:char(24)"`
}
