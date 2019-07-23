//+build !test

package utils

import "github.com/gin-gonic/gin"

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

/*
GetQueryFile get query file info
return file byte, response code and error
response code could be 200 or 413 or 500
*/
func GetQueryFile(c *gin.Context, name string, maxsize int64) ([]byte, int, error) {
	if CheckInTest() {
		return []byte("test_file"), 200, nil
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
	return nil, 500, err
}
