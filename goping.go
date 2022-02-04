// Script to query url N times simultaneously
package main

import (
	"os"
	"strconv"

	"github.com/dd84ai/goping/core"
)

func main() {
	channels_amount, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	url := os.Args[2]

	core.Queries(channels_amount, url)
}
