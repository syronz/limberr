package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/syronz/limberr"
)

func main() {
	fmt.Println("vim-go")

	err := errors.New("new_error")
	err = limberr.AddCode(err, "E10")
	err = limberr.AddMessage(err, "a_message")
	err = limberr.AddType(err, "regular_teype", "title")
	err = limberr.AddPath(err, "/path/to/somewhere")
	err = limberr.AddStatus(err, 403)
	err = limberr.AddDomain(err, "base")

	t.Log("Err: ", err)

	jsonErr, _ := json.MarshalIndent(err, "", "    ")
	// func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	t.Log("jsonErr: ", string(jsonErr))

	translator := func(str string, a ...interface{}) string {
		return str
	}

	_ = translator

	statusCode, parsedErr := limberr.Parse(err, translator)

	t.Log("statusCode: ", statusCode)
	t.Log("parsed error: ", parsedErr)
	jsonParsedErr, _ := json.MarshalIndent(parsedErr, "", "    ")
	t.Log("statusCode: ", string(jsonParsedErr))
}
