package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if  name=="" {
		return "", errors.New("No name provided")
	}
	message := fmt.Sprintf("Hi, welcome %v!", name)
	return message, nil
}
