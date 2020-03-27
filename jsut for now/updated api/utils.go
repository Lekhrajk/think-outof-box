package main

import (
	"fmt"
	"os"
	"regexp"
)

// Regex to extract the error message from OpenNebula response
var errmsgRegex = regexp.MustCompile(`.* \[.*\] (.*)`)

// Used for user defined errors
type errorJSON struct {
	Error string `json:"error"`
}

// Used for user defined errors
type messageJSON struct {
	Message string `json:"message"`
}

// Get the SECRET KEY for encrypting and decrypting JWT
func getSecretKey() []byte {
	jwtKey := []byte(os.Getenv("XONGL_JWT_SECRET"))
	return jwtKey
}

// Used for OpenNebula  errors response
func errorMsg(err error) errorJSON {
	msg := errmsgRegex.FindStringSubmatch(err.Error())
	return errorJSON{
		Error: msg[1],
	}
}

// Used for providing user defined message response
func okMsg(format string, i ...interface{}) messageJSON {
	return messageJSON{
		Message: fmt.Sprintf(format, i...),
	}
}
