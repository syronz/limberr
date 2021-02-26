# limberr
limberr is a Go package for handling errors. It supports (RFC-7807)[https://tools.ietf.org/html/rfc7807].

[![Go Report Card](https://goreportcard.com/badge/github.com/syronz/limberr)](https://goreportcard.com/report/github.com/syronz/limberr)

limberr implements error handling based on chain and parent/child
relation, id consist two part. generating the error and parse it.

## Generation Error
there is two way for using this package simple way is using direct method

### Generate Error by Method
```go
err := errors.New("new_error")
err = limberr.AddCode(err, "E10")
err = limberr.AddMessage(err, "a_message")
err = limberr.AddType(err, "regular_type", "title")
err = limberr.AddPath(err, "/path/to/somewhere")
err = limberr.AddStatus(err, 403)
err = limberr.AddDomain(err, "base")

translator := func(str string, a ...interface{}) string {
  return str
}

statusCode, parsedErr := limberr.Parse(err, translator)
```

Output be like below
```JSON
statusCode:  403
statusCode:  {
    "code": "E10",
    "type": "regular_type",
    "title": "title",
    "domain": "base",
    "message": "a_message",
    "path": "/path/to/somewhere",
    "original_error": "new_error"
}
```

### Generate Error by Builder
Builder is more complicated but at the end need less code, you can find the complete instance inside the [examples](https://github.com/syronz/limberr/tree/main/examples) directory
```go
err := limberr.New("new_error", "E10").
  Message("a_message %v %v", "param1", "param2").
  Custom(ForbiddenErr).
  Path("/path/to/somewhere").
  Build()
```
output is as same as before


### invalid params
easily by adding invalid params you can have perfect error for highligt fields which have error with
proper message
```go
InvalidParam("age", "age should be more than %v", 18).
```

## Parsing
For parsing you can use it in response to add path in easiear way

Final format be something like this
```json
{
  "type": "http//link.com/to/order",
  "title": "duplication",
  "message": "user with this name already exist",
  "code": "E321343",
  "path": "users/32",
  "invalid-params": [
    {
      "name": "age",
      "reason": "must be a positive integer"
    },
    {
      "name": "color",
      "reason": "must be 'green', 'red' or 'blue'"
    }
  ]
}
```

## Dictionary
By combining this package with dict and create a manual function as translator you can have multi
language error generator
[github.com/syronz/dict](https://github.com/syronz/dict)
