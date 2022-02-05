// Script to query url N times simultaneously
package main

import (
	"flag"
	"fmt"

	"goping/query"
	"os"
)

func main() {
	fmt.Println(os.Args)
	channels_amount := flag.Int("n", 1, "Amount times to query the url")
	url := flag.String("url", "", "Url to query")
	flag.Parse()

	fmt.Println(*channels_amount)
	fmt.Println(*url)

	fmt.Println("parsed arguments")

	query.Queries(*channels_amount, *url)
}
