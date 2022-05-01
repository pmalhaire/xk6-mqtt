package mqtt

import (
	"errors"
)

var (
	ErrorConnect      = errors.New("connection failed")
	ErrorState        = errors.New("invalid state")
	ErrorClient       = errors.New("invalid client given")
	ErrorTimeout      = errors.New("operation timeout")
	ErrorSubscribe    = errors.New("subscribe failure")
	ErrorConsumeToken = errors.New("invalid consume token")
	ErrorPublish      = errors.New("publish failure")
)
