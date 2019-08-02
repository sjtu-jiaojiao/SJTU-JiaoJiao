package db

import (
	"time"
)

// Test database map
type Test struct {
	ID       int32  `gorm:"auto_increment;primary_key"`
	TestName string `gorm:"size:100"`
}

// User database map
type User struct {
	ID          int32  `gorm:"auto_increment;primary_key"`
	UserName    string `gorm:"size:32;not null"`
	AvatarID    string `gorm:"type:char(24)"`
	Telephone   string `gorm:"type:char(11)"`
	StudentID   string `gorm:"size:32;index;not null"`
	StudentName string `gorm:"size:32;not null"`
	Status      int32  `gorm:"default:1;not null"`
	Role        int32  `gorm:"default:1;not null"`
}

// SellInfo database map
type SellInfo struct {
	ID          int32     `gorm:"auto_increment;primary_key"`
	Status      int32     `gorm:"default:1;not null"`
	ReleaseTime time.Time `gorm:"not null"`
	ValidTime   time.Time `gorm:"not null"`
	UserID      int32     `gorm:"not null;index"`
	GoodID      int32     `gorm:"not null;index"`
}

// BuyInfo database map
type BuyInfo struct {
	ID          int32     `gorm:"auto_increment;primary_key"`
	Status      int32     `gorm:"default:1;not null"`
	ReleaseTime time.Time `gorm:"not null"`
	ValidTime   time.Time `gorm:"not null"`
	UserID      int32     `gorm:"not null;index"`
	GoodID      int32     `gorm:"not null;index"`
}

//Transaction database map
type Transaction struct {
	ID         int32     `gorm:"auto_increment;primary_key"`
	InfoID     int32     `gorm:"not null"`
	Category   int32     `gorm:"not null"`
	FromUserID int32     `gorm:"not null"`
	ToUserID   int32     `gorm:"not null"`
	CreateTime time.Time `gorm:"not null"`
	Status     int32     `gorm:"not null;default:1"`
}

// Good database map
type Good struct {
	ID          int32   `gorm:"auto_increment;primary_key"`
	GoodName    string  `gorm:"size:128;not null"`
	Price       float64 `gorm:"default:0;not null"`
	Description string  `gorm:"type:text"`
	ContentID   string  `gorm:"type:char(24)"`
}
