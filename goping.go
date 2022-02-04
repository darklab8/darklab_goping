// Script to query url N times simultaneously
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

func query(answer chan string, url string) {

	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("query failed")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("readAll failed")
	}

	elapsed := time.Since(start)
	answer <- string(fmt.Sprintf(string(body), elapsed))
}

func main() {

	start := time.Now()

	channels_amount, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	chans := []chan string{}
	for i := 0; i < channels_amount; i++ {
		chans = append(chans, make(chan string))
	}

	url := os.Args[2]

	for _, channel := range chans {
		go query(channel, url)
	}

	for _, channel := range chans {
		println(<-channel)
	}

	elapsed := time.Since(start)
	fmt.Printf("total time %s\n", elapsed)
}
