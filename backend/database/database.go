package db

import (
	"time"
)

// Test is db test table map
type Test struct {
	Id       int32  `orm:"auto;pk;column(test_id)"`
	TestName string `orm:"size(100)"`
}

// User is db user table map
type User struct {
	Id          int32  `orm:"auto;pk;column(user_id)"`
	UserName    string `orm:"size(32)"`
	AvatarId    string `orm:"type(char);size(24)"`
	Telephone   string `orm:"type(char);size(11);null"`
	StudentId   string `orm:"size(32);index"`
	StudentName string `orm:"size(32)"`
	Status      int32
}

// SellInfo is db release table map
type SellInfo struct {
	Id          int32 `orm:"auto;pk;column(sell_info_id)"`
	Status      int32
	ReleaseTime time.Time `orm:"auto_now_add;type(datetime)"`
	ValidTime   time.Time `orm:"type(datetime)"`
	UserId      int32     `orm:"index"`
	GoodId      int32     `orm:"index"`
}

// Good is db good table map
type Good struct {
	Id          int32  `orm:"auto;pk;column(good_id)"`
	GoodName    string `orm:"size(128)"`
	Price       float64
	Description string `orm:"type(text);null"`
	ContentId   string `orm:"type(char);size(24)"`
}
