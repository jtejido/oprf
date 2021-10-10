package oprf

import (
	"errors"
)

var (
	ErrI2OSPLargeInteger    = errors.New("Integer too large")
	ErrI2OSPNegativeInteger = errors.New("Negative integer")
)
