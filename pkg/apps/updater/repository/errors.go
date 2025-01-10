package repository

import "fmt"

type errorItemNotChanged struct {
	identifier string
}

func (e *errorItemNotChanged) Error() string {
	return fmt.Sprintf("item '%s' has not being changed", e.identifier)
}

func NewErrorItemNotChanged(identifier string) error {
	return &errorItemNotChanged{
		identifier: identifier,
	}
}

func IsErrorItemNotChanged(e error) (string, bool) {
	if re, ok := e.(*errorItemNotChanged); ok {
		return re.identifier, true
	}
	return "", false
}
