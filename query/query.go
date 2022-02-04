package query

import (
	"fmt"
	"io/ioutil"
	"net/http"
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

func Queries(channels_amount int, url string) {
	start := time.Now()

	chans := []chan string{}
	for i := 0; i < channels_amount; i++ {
		chans = append(chans, make(chan string))
	}

	for _, channel := range chans {
		go query(channel, url)
	}

	for _, channel := range chans {
		println(<-channel)
	}

	elapsed := time.Since(start)
	fmt.Printf("total time %s\n", elapsed)
}
