package telnetClient

import (
	"errors"
	"strings"
)

// parses string response
func ParseResponse(res string) ([]string, error) {
	r := strings.Split(res, ", ")
	if r[0] == "err" {
		return make([]string, 0), errors.New("Response from server: " + strings.Join(r[1:], ", "))
	}

	return r[1:], nil
}
