package badlink_test

import (
	"errors"
	"testing"

	"."
)

type someType struct{}

func (s *someType) MarshalYAML() (interface{}, error) {
	return nil, errors.New("thing")
}

func TestThing(t *testing.T) {
	badlink.Thing(&someType{})
}
