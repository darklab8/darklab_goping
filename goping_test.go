package main

import (
	"fmt"
	"goping/query"
	"goping/webserver"
	"os/exec"
	"strings"
	"testing"
)

func TestTryingSimpleCMD(t *testing.T) {
	cmd := exec.Command("echo", "123")
	stdout, err := cmd.Output()

	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(stdout))
}

func TestCMD(t *testing.T) {
	webserver.Activate_test_web_server(8081)

	// Internal
	body := query.Queries(1, "http://localhost:8081/")

	fmt.Println(body)
	if strings.Count(body, "pong") != 1 {
		t.Error("I did not get pong!")
	}

	// CMD
	cmd := exec.Command("go", "run", ".", "-url=http://localhost:8081/", "-n=2")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(string(stdout))
		t.Error(err)
	}

	fmt.Println(string(stdout))
}
