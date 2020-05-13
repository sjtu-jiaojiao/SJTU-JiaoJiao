//+build !test

package utils

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func check(actual interface{}, expect interface{}) {
	switch actual := actual.(type) {
	case string:
		t := expect.(string)
		if t == "#NOTEMPTY#" {
			So(actual, ShouldNotBeEmpty)
		} else {
			So(t, ShouldEqual, actual)
		}
	case map[string]interface{}:
		t := expect.(map[string]interface{})
		for k, v := range actual {
			check(v, t[k])
		}
	case []interface{}:
		t := expect.([]interface{})
		for k, v := range actual {
			check(v, t[k])
		}
	default:
		switch expect := expect.(type) {
		case string:
			if expect == "#NOTZERO#" {
				So(actual, ShouldNotBeZeroValue)
			}
		default:
			So(expect.(float64), ShouldEqual, actual)
		}
	}
}

func TransVar(s string, verify map[string]interface{}, variable map[string]interface{}) interface{} {
	v := verify[s].(string)
	if string(v[0]) == "#" && string(v[len(v)-1]) == "#" {
		return variable[v[1:len(v)-1]]
	}
	return verify[s]
}

func Test(t *testing.T, file string, data func(param interface{}),
	parse func(param interface{}) map[string]interface{},
	verify func(param interface{}, output interface{})) {
	Convey("Test "+file, t, func() {
		fileData, err := ioutil.ReadFile(file)
		So(err, ShouldBeNil)

		var fileMap map[string][]map[string]interface{}
		So(json.Unmarshal(fileData, &fileMap), ShouldBeNil)
		if data != nil {
			for _, v := range fileMap["data"] {
				data(v)
			}
		}

		for _, v := range fileMap["case"] {
			ret := parse(v["input"])
			check(ret, v["output"])
			if tmp, ok := v["verify"]; ok {
				switch tmp := tmp.(type) {
				case []interface{}:
					for _, vtmp := range tmp {
						verify(vtmp, ret)
					}
				default:
					verify(v["verify"], ret)
				}
			}
		}
	})
}
