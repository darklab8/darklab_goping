package query

import (
	"fmt"
	"goping/webserver"
	"strings"
	"testing"
)

func TestQueryingAccessToWebserverPackage(t *testing.T) {
	msg := webserver.Ping()
	fmt.Println(msg)

	if msg != "pong!" {
		t.Error("not a pong answer")
	}
}

func TestWebServerForOneRequest(t *testing.T) {
	webserver.Activate_test_web_server()
	body := query("http://localhost:8080/")

	fmt.Println(body)

	if !strings.Contains(body, "pong!") {
		t.Error("I did not get pong!")
	}
}

func TestWebServerForMultipleRequests(t *testing.T) {
	webserver.Activate_test_web_server()
	body := Queries(5, "http://localhost:8080/")

	fmt.Println(body)
	if strings.Count(body, "pong!") != 5 {
		t.Error("I did not get pong!")
	}
}
