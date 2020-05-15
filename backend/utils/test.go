//+build !test

package utils

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type StringMap = map[string]interface{}

func TestCheck(actual interface{}, expect interface{}) {
	switch actual := actual.(type) {
	case string:
		t := expect.(string)
		if t == "#NOTEMPTY#" {
			So(actual, ShouldNotBeEmpty)
		} else if strings.HasPrefix(t, "#NOT#") {
			So(actual, ShouldNotEqual, t[5:])
		} else {
			So(actual, ShouldEqual, t)
		}
	case map[string]interface{}:
		t := expect.(map[string]interface{})
		So(len(t), ShouldBeLessThanOrEqualTo, len(actual))
		for k, v := range actual {
			if k != "_error" {
				TestCheck(v, t[k])
			}
		}
	case []interface{}:
		t := expect.([]interface{})
		So(len(t), ShouldBeLessThanOrEqualTo, len(actual))
		for k, v := range actual {
			TestCheck(v, t[k])
		}
	default:
		switch expect := expect.(type) {
		case string:
			if expect == "#NOTZERO#" {
				So(actual, ShouldNotBeZeroValue)
			} else if strings.HasPrefix(expect, "#NOT#") {
				num, err := strconv.Atoi(expect[5:])
				So(err, ShouldBeNil)
				So(actual, ShouldNotEqual, num)
			} else {
				LogPanic("Output type differ")
			}
		default:
			So(actual, ShouldEqual, expect)
		}
	}
}

func TransVar(s string, verify StringMap, variable StringMap) interface{} {
	switch v := verify[s].(type) {
	case string:
		if string(v[0]) == "#" && string(v[len(v)-1]) == "#" {
			return variable[v[1:len(v)-1]]
		}
	}
	return verify[s]
}

func TransVarInt(s string, verify StringMap, variable StringMap) int32 {
	return TestInt(TransVar(s, verify, variable))
}

func TransVarString(s string, verify StringMap, variable StringMap) string {
	return TestString(TransVar(s, verify, variable))
}

func TestInt(d interface{}) int32 {
	return int32(d.(float64))
}

func TestInt64(d interface{}) int64 {
	return int64(d.(float64))
}

func TestUint(d interface{}) uint32 {
	return uint32(d.(float64))
}

func TestString(d interface{}) string {
	return d.(string)
}

func TestByte(d interface{}) []byte {
	return []byte(d.(string))
}

func TestBool(d interface{}) bool {
	return d.(bool)
}

func TestFloat64(d interface{}) float64 {
	return d.(float64)
}

func Test(t *testing.T, file string,
	insert func(data StringMap),
	parse func(input StringMap) StringMap,
	verify func(verify StringMap, output StringMap),
	checker func(actual StringMap, expect StringMap)) {
	Convey("Test "+file, t, func() {
		t.Logf("Reading file %s...", file)
		fileData, err := ioutil.ReadFile(file)
		So(err, ShouldBeNil)

		t.Log("Init data...")
		var fileMap map[string][]map[string]interface{}
		So(json.Unmarshal(fileData, &fileMap), ShouldBeNil)
		if insert != nil {
			for _, v := range fileMap["data"] {
				insert(v)
			}
		}

		for i, v := range fileMap["case"] {
			t.Logf("Testing case %d...", i)
			ret := parse(v["input"].(StringMap))
			err, ok := v["output"].(map[string]interface{})["_error"]
			if ok && err.(bool) {
				So(ret["_error"], ShouldNotBeNil)
			} else {
				So(ret["_error"], ShouldBeNil)
				if checker != nil {
					checker(ret, v["output"].(StringMap))
				} else {
					TestCheck(ret, v["output"])
				}
			}
			if tmp, ok := v["verify"]; ok {
				switch tmp := tmp.(type) {
				case []interface{}:
					for _, vtmp := range tmp {
						verify(vtmp.(StringMap), ret)
					}
				default:
					verify(v["verify"].(StringMap), ret)
				}
			}
			t.Logf("Case %d OK", i)
		}
	})
}
