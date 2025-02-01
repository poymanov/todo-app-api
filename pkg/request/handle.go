package request

import (
	"errors"
	"net/http"
	"reflect"
)

func HandleBody[T any](request *http.Request) (*T, error) {
	body, err := decode[T](request.Body)

	if reflect.ValueOf(body).IsZero() {
		return nil, errors.New("wrong request body")
	}

	if err != nil {
		return nil, err
	}

	err = isValid(body)

	if err != nil {
		return nil, err
	}

	return &body, nil
}
