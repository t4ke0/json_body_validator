# json_request_validator




## Example


```golang
package main

import (
    "fmt"

    "github.com/t4ke0/json_request_validator"
)

// EchoRequest
type EchoRequest struct {
	Message  string `json:"message" validator:"required"`
	Username string `json:"username" validator:"optional"`
	Data     struct {
		FullName string `json:"full_name" validator:"required"`
	} `json:"data" validator:"required"`
}

var port string = os.Getenv("PORT")

func main() {

	var req EchoRequest = EchoRequest{
		Message: "some message here",
		Data: struct {
			FullName string `json:"full_name" validator:"required"`
		}{
			FullName: "user",
		},
	}

	fmt.Println("JSON request is valid ?", json_request_validator.Validate(req))
}
```
