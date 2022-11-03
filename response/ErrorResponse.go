package response

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

type Response struct {
	Key     string `json:"key"`
	Message string `json:"message"`
}

var repo map[string]interface{}

func init() {
	repo = readErrorFile()
}

func GetErrorResponse(name string) string {
	return repo[name].(string)
}

func GetResponseByKey(key string) *Response {
	return &Response{
		Key:     key,
		Message: GetErrorResponse(key),
	}
}

func readErrorFile() map[string]interface{} {
	byteValue, err := os.ReadFile("response/errorResponse.json")

	if err != nil {
		log.Fatal(err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(byteValue, &result)

	if err != nil {
		log.Fatal(err)
	}

	return result
}

var (
	ErrInvalidId       error = errors.New("invalid_id")
	ErrDuplicate       error = errors.New("already_exists")
	ErrUnexpected      error = errors.New("unexpected_error")
	ErrInvalidStatus   error = errors.New("invalid_status")
	ErrInvalidUsername error = errors.New("invalid_username")
	ErrInvalidPassword error = errors.New("invalid_password")
)
