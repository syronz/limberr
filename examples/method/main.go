package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/syronz/limberr"
)

func main() {

	err := errors.New("new_error")
	err = limberr.AddCode(err, "E10")
	err = limberr.AddMessage(err, "a_message")
	err = limberr.AddType(err, "regular_teype", "title")
	err = limberr.AddPath(err, "/path/to/somewhere")
	err = limberr.AddStatus(err, 403)
	err = limberr.AddDomain(err, "base")

	translator := func(str string, a ...interface{}) string {
		return str
	}

	statusCode, parsedErr := limberr.Parse(err, translator)

	fmt.Println("parsed error: ", parsedErr)
	fmt.Println("statusCode: ", statusCode)
	jsonParsedErr, _ := json.MarshalIndent(parsedErr, "", "    ")
	fmt.Println("error: ", string(jsonParsedErr))
}
