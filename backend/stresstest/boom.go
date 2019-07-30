package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/myzhan/boomer"
)

var names []string
var rnd *rand.Rand

func goodSearch() {
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	start := time.Now()
	res, err := httpClient.Get("http://202.120.40.8:30711/v1/sellInfo?" + url.Values{
		"limit":    {"20"},
		"goodName": {names[rnd.Intn(len(names))]},
	}.Encode())
	if err != nil {
		boomer.RecordFailure("http", "goodSearch", time.Since(start).Nanoseconds()/int64(time.Millisecond), err.Error())
		return
	}
	body, err := ioutil.ReadAll(res.Body)
	elapsed := time.Since(start)

	if err != nil {
		boomer.RecordFailure("http", "goodSearch", elapsed.Nanoseconds()/int64(time.Millisecond), err.Error())
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		boomer.RecordFailure("http", "goodSearch", elapsed.Nanoseconds()/int64(time.Millisecond), res.Status)
		return
	}
	boomer.RecordSuccess("http", "goodSearch", elapsed.Nanoseconds()/int64(time.Millisecond), int64(len(body)))
}

func main() {
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

	fp, err := os.Open("goodbase.txt")
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	bufReader := bufio.NewReader(fp)

	for {
		line, _, err := bufReader.ReadLine()
		if err != nil {
			if err == io.EOF {
				err = nil
				break
			}
		} else {
			names = append(names, string(line))
		}
	}

	task1 := &boomer.Task{
		Name:   "goodSearch",
		Weight: 10,
		Fn:     goodSearch,
	}

	boomer.Run(task1)
}
