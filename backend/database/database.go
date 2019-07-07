package db

// Test is db test table map
type Test struct {
	Id       int    `orm:"auto;pk;column(test_id)"`
	TestName string `orm:"size(100)"`
}

// User is db user table map
type User struct {
	Id          int `orm:"auto;pk;column(user_id)"`
	StudentId   uint64
	StudentName string `orm:"size(100)"`
}
