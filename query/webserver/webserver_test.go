package webserver

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestEncodingAnswer(t *testing.T) {
	message := map[string]string{"msg": "pong!"}
	encoded_message, err := json.Marshal(message)

	if err != nil {
		log.Fatal(err)
	}

	encoded_message2 := []byte("{\"msg\":\"pong!\"}")

	fmt.Println(string(encoded_message))
	fmt.Println(string(encoded_message2))

	string1 := string(encoded_message)
	string2 := string(encoded_message2)
	if string1 != string2 {
		t.Errorf("%v, %v aren\t equal by value", string1, string2)
	}

	if reflect.TypeOf(encoded_message) != reflect.TypeOf(encoded_message2) {
		t.Errorf("%v, %v aren\t equal by type", string1, string2)
	}
}
