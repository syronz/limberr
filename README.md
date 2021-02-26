# limberr
limberr is a Go package for handling errors. It supports RFC 7807

## usage 
there is two way for using this package simple way is using direct method

### method
```
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
```

the output be like below
```
statusCode:  403
statusCode:  {
    "code": "E10",
    "type": "regular_teype",
    "title": "title",
    "domain": "base",
    "message": "a_message",
    "path": "/path/to/somewhere",
    "original_error": "new_error"
}
```

### builder
builder is simpler and need less code, you can find the complete instance inside the examples directory
```
	err := limberr.New("new_error", "E10").
		Message("a_message %v %v", "param1", "param2").
		Custom(ForbiddenErr).
		Path("/path/to/somewhere").
		Build()
```
the output be as same as before


### invalid params
easily by adding invalid params you can have perfect error for highligt fields which have error with
proper message
```
		InvalidParam("age", "age should be more than %v", 18).
```
