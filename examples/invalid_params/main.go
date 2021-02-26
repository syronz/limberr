package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/syronz/limberr"
)

const (
	Nil limberr.CustomError = iota
	ForbiddenErr
)

var UniqErrorMap limberr.CustomErrorMap

func main() {
	UniqErrorMap = make(map[limberr.CustomError]limberr.ErrorTheme)
	UniqErrorMap[ForbiddenErr] = limberr.ErrorTheme{
		Type:   "#FORBIDDEN",
		Title:  "title",
		Domain: "base",
		Status: http.StatusConflict,
	}

	err := limberr.New("new_error", "E10").
		Message("a_message %v %v", "param1", "param2").
		Custom(ForbiddenErr).
		Path("/path/to/somewhere").
		InvalidParam("name", "name is required").
		InvalidParam("age", "age should be more than %v", 18).
		Build()

	// you can define translator to get the params for message, so the flexible
	// messages are available in multi language also
	translator := func(str string, a ...interface{}) string {
		return str
	}

	customError := limberr.GetCustom(err)
	err = limberr.ApplyCustom(err, UniqErrorMap[customError], "path/to/explain/error")
	statusCode, parsedErr := limberr.Parse(err, translator)

	// result
	fmt.Println("log parsed error: ", parsedErr)
	fmt.Println("statusCode: ", statusCode)
	jsonParsedErr, _ := json.MarshalIndent(parsedErr, "", "    ")
	fmt.Println("error: ", string(jsonParsedErr))
}
