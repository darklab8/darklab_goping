// Script to query url N times simultaneously
package main

import (
	"os"
	"strconv"

	"github.com/dd84ai/goping/query"
)

func main() {
	channels_amount, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	url := os.Args[2]

	query.Queries(channels_amount, url)
}
