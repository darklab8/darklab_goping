// package purely for testing purposes
package webserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func viewHandler(writer http.ResponseWriter, request *http.Request) {

	message := map[string]string{"msg": "pong!"}
	encoded_message, err := json.Marshal(message)

	if err != nil {
		log.Fatal(err)
	}

	_, err = writer.Write(encoded_message)
	if err != nil {
		log.Fatal(err)
	}
}

func Webserver(port int) {
	http.HandleFunc("/", viewHandler)
	err := http.ListenAndServe(fmt.Sprintf("localhost:%d", port), nil)
	log.Fatal(err)
}

var (
	is_webserver_active bool = false
)

func Activate_test_web_server(port_optional ...int) {

	port := 8080
	if len(port_optional) > 0 {
		port = port_optional[0]
	}

	if !is_webserver_active {
		is_webserver_active = true
		go Webserver(port)
		time.Sleep(200 * time.Millisecond)
	}
}

func Ping() string {
	return "pong!"
}
