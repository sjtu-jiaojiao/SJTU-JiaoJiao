//+build !test

package utils

import (
	"bytes"
	"fmt"

	"github.com/gin-gonic/gin"
)

// AssignNotEmpty assign a string when another is not empty
func AssignNotEmpty(src *string, dst *string) {
	if *src != "" {
		*dst = *src
	}
}

// AssignNotZero assign a digit when another is not zero
func AssignNotZero(src interface{}, dst interface{}) {
	switch v := dst.(type) {
	case *int32:
		if *src.(*int32) != 0 {
			*v = *src.(*int32)
		}
	case *float32:
		if *src.(*float32) != 0 {
			*v = *src.(*float32)
		}
	case *float64:
		if *src.(*float64) != 0 {
			*v = *src.(*float64)
		}
	default:
		panic("wrong type")
	}
}

// IsEmpty check param zero value
func IsEmpty(val interface{}) bool {
	switch v := val.(type) {
	case int:
		return v == 0
	case int32:
		return v == 0
	case int64:
		return v == 0
	case float32:
		return v == 0
	case float64:
		return v == 0
	case string:
		return v == ""
	case []byte:
		return bytes.Equal(v, []byte{0}) || string(v) == ""
	case []string:
		return len(v) == 0
	default:
		s := fmt.Sprintf("%d", val)
		return s == "0"
	}
}

// RequireParam check param and return true for all param provided
func RequireParam(v ...interface{}) bool {
	for _, arg := range v {
		if IsEmpty(arg) {
			return false
		}
	}
	return true
}

/*
GetQueryFile get query file info
return file byte, response code and error
response code could be 200 or 413 or 500
*/
func GetQueryFile(c *gin.Context, name string, maxsize int64) ([]byte, int, error) {
	if CheckInTest() {
		s := c.PostForm(name)
		if s != "" {
			return []byte(s), 200, nil
		} else {
			return nil, 400, nil
		}
	}

	file, err := c.FormFile(name)
	if err == nil {
		if file.Size > maxsize {
			return nil, 413, nil
		}

		f, err := file.Open()
		if LogContinue(err, Warning) {
			return nil, 500, err
		}
		defer f.Close()
		data := make([]byte, file.Size)
		_, err = f.Read(data)
		if LogContinue(err, Warning) {
			return nil, 500, err
		}
		return data, 200, nil
	}
	return nil, 400, nil
}

// EnumConvert convert int32 to enum
func EnumConvert(v int32, n map[int32]string) int32 {
	if n[v] != "" {
		return v
	}
	return 0
}

// CheckFile check if file valid
func CheckFile(file []byte, f ...func(buf []byte) bool) bool {
	if CheckInTest() {
		return string(file) == "valid_file"
	}
	for _, v := range f {
		if v(file) {
			return true
		}
	}
	return false
}
