package hipchat

import (
	"errors"
	"fmt"
)

var ErrUnknown = errors.New("An unknown error occurred.")

type Error struct {
	Code    float64 `json:"code"`
	Message string  `json:"message"`
	Type    string  `json:"type"`
}

func (e Error) Error() string {
	// TODO(tmc): improve formatting
	return fmt.Sprintf("%v: %v", e.Code, e.Type)
}
