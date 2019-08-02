package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/myzhan/boomer"
)

var names []string
var rnd *rand.Rand
var t *http.Transport
var c *http.Client

func goodSearch() {
	start := time.Now()
	//res, err := httpClient.Get("http://202.120.40.8:30711/v1/sellInfo?" + url.Values{
	//	"limit":    {"20"},
	//	"goodName": {names[rnd.Intn(len(names))]},
	//}.Encode())
	res, err := c.Get("http://202.120.40.8:30711/v1/user/" + fmt.Sprintf("%v", rnd.Intn(9999)+1))
	if err != nil {
		boomer.RecordFailure("http", "goodSearch", time.Since(start).Nanoseconds()/int64(time.Millisecond), err.Error())
		return
	}
	length, err := io.Copy(ioutil.Discard, res.Body)
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
	boomer.RecordSuccess("http", "goodSearch", elapsed.Nanoseconds()/int64(time.Millisecond), length)
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

	t = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		MaxIdleConns:        10240,
		MaxIdleConnsPerHost: 10240,
		DisableKeepAlives:   true,
	}
	c = &http.Client{
		Timeout:   40 * time.Second,
		Transport: t,
	}

	task1 := &boomer.Task{
		Name:   "goodSearch",
		Weight: 10,
		Fn:     goodSearch,
	}

	boomer.Run(task1)
}
