package query

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func query(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("query failed")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("readAll failed")
	}

	return string(body)
}

func channeledQuery(answer chan string, url string) {

	start := time.Now()

	body := query(url)

	elapsed := time.Since(start)
	answer <- string(fmt.Sprintf(body, elapsed))
}

func Queries(channels_amount int, url string) string {
	start := time.Now()

	chans := []chan string{}
	for i := 0; i < channels_amount; i++ {
		chans = append(chans, make(chan string))
	}

	for _, channel := range chans {
		go channeledQuery(channel, url)
	}

	var sb strings.Builder
	for _, channel := range chans {
		msg := <-channel
		fmt.Println(msg)
		sb.WriteString(msg)
	}

	elapsed := time.Since(start)
	fmt.Printf("total time %s\n", elapsed)

	return sb.String()
}
