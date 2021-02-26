package limberr

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestParse(t *testing.T) {
	err := errors.New("new_error")
	err = AddCode(err, "E10")
	err = AddMessage(err, "a_message")
	err = AddType(err, "regular_teype", "title")
	err = AddPath(err, "/path/to/somewhere")
	err = AddStatus(err, 403)
	err = AddDomain(err, "base")

	t.Log("Err: ", err)

	jsonErr, _ := json.MarshalIndent(err, "", "    ")
	// func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	t.Log("jsonErr: ", string(jsonErr))

	translator := func(str string, a ...interface{}) string {
		return str
	}

	_ = translator

	statusCode, parsedErr := Parse(err, translator)

	t.Log("statusCode: ", statusCode)
	t.Log("parsed error: ", parsedErr)
	jsonParsedErr, _ := json.MarshalIndent(parsedErr, "", "    ")
	t.Log("statusCode: ", string(jsonParsedErr))

}
