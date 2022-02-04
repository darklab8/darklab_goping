// Script to query url N times simultaneously
package main

import (
	"flag"

	"github.com/dd84ai/goping/query"
)

func main() {
	channels_amount := flag.Int("n", 1, "Amount times to query the url")
	url := flag.String("url", "", "Url to query")
	flag.Parse()

	query.Queries(*channels_amount, *url)
}
