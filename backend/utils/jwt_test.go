package utils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

/*
alg="HS256"
secret="test_secret"
*/

func TestJWTVerify(t *testing.T) {
	tf := func(token string) error {
		_, err := JWTVerify(token, "test_secret")
		return err
	}
	Convey("JWTVerify test", t, func() {
		// normal
		So(tf("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZXN0IjoidGVzdF9wYXlsb2FkIn0.gzg19T-vgVjdSskOjwXJnKxDjO8hp3bXG_PS6iFwKts"), ShouldEqual, nil)
		// expired
		So(tf("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZXN0IjoidGVzdF9wYXlsb2FkIiwiZXhwIjoxNTQ2MjcyMDAwfQ._1yb5UkyPYJqSfJ5y9pKsB2kbr_EpibVhKa2RvUl68A"), ShouldBeError, "Token is expired")
		// not expired
		So(tf("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZXN0IjoidGVzdF9wYXlsb2FkIiwiZXhwIjo0NzAxOTQ1NjAwfQ.5fZOk3TLL6h3vGD7W0i3foRZeGn5eaLJxsymUznp4zw"), ShouldEqual, nil)
		// invalid
		So(tf("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZXN0IjoidGVzdF9wYXlsb2FkIn0.UiG1YraTzdgbq--sxTV8lsJMr4sfZeVkB-KAyKSS3bU"), ShouldBeError, "signature is invalid")
		// invalid
		So(tf("eyJhbGciOiJIUzI1NiIsnR5cCI6IkpXVJ9.eyJ0ZXN0IjoidGVzdF9wYXlsb2FkIn0.UiG1YraTdgbq--sxTV8lsJMr4sfZVkB-KAyKS3bU"), ShouldBeError)
		// not support
		So(tf("eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZXN0IjoidGVzdF9wYXlsb2FkIn0.rfsMIWo6HwV79Xiq53RnTeylFsukuXv0WF3JJ9_wBu-BEMbujaaQW2rvKfPthlVMQuubkcD2ENZdp7ZNIlAE3Q"), ShouldBeError, "Unexpected signing method: ES256")
	})
}

func TestJWTParse(t *testing.T) {
	Convey("JWTParse test", t, func() {
		// valid
		t, err := JWTVerify("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZXN0IjoidGVzdF9wYXlsb2FkIiwiZXhwIjo0NzAxOTQ1NjAwfQ.5fZOk3TLL6h3vGD7W0i3foRZeGn5eaLJxsymUznp4zw", "test_secret")
		So(err, ShouldEqual, nil)
		So(JWTParse(t, "test"), ShouldEqual, "test_payload")
		So(JWTParse(t, "test1"), ShouldEqual, "")
	})
}
