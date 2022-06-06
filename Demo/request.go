package main_demo

import (
	"fmt"
	"net/http"
)

func download_demo() {
	req, err := http.NewRequest("GET", "https://www.apache.org/dyn/closer.lua/zookeeper/zookeeper-3.8.0/apache-zookeeper-3.8.0-bin.tar.gz", nil)
	if err != nil {
		println(err)
	}
	range_begin := 2000
	range_end := 3000
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", range_begin, range_end))
	fmt.Println(req.Header)
	println(">>>>>><<<<<<")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		println(err)
	}
	// when will we be able to use ↓
	// res:=, err= http.DefaultClient.Do(req)  ?
	fmt.Println(res)
}
