// Generated file, please do not change!!!
package ml

import (
	"fmt"
)

type GenericRequestError struct {
	Content    []byte
	StatusCode int
}

func (e GenericRequestError) Error() string {
	return fmt.Sprintf("Request returned status code %d", e.StatusCode)
}
