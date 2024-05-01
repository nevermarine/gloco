package errdefs

import "errors"

var (
	ErrMultipleServices = errors.New("multiple services in Compose file")
	ErrNoService = errors.New("no service in Compose file")
)

func ContainsMultipleServices(err error) bool {
	return errors.Is(err, ErrMultipleServices)
}

func ContainsNoService(err error) bool {
	return errors.Is(err, ErrNoService)
}