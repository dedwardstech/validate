package validate

import "github.com/pkg/errors"

// Fn ...
type Fn func(interface{}) error

// ErrInvalidType ...
var ErrInvalidType = errors.New("invalid type")
