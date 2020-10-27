package validate

import "github.com/pkg/errors"

// MinLength validates that the length of a string or array is greater than
// or equal to the minLength value passed in
func MinLength(minLength int, errMsg string) Fn {
	err := errors.New(errMsg)

	return func(item interface{}) error {
		r := compareLength(item, minLength)

		if r == uncomparable {
			return ErrInvalidType
		}

		if r == lessThan {
			return err
		}

		return nil
	}
}

// MaxLength validates that the length of a string or array is greater than
// or equal to the minLength value passed in
func MaxLength(maxLength int, errMsg string) Fn {
	err := errors.New(errMsg)

	return func(item interface{}) error {
		r := compareLength(item, maxLength)

		if r == uncomparable {
			return ErrInvalidType
		}

		if r == greaterThan {
			return err
		}

		return nil
	}
}

// Between ...
func Between(a, b int, errMsg string) Fn {
	betweenErr := errors.New(errMsg)

	return func(item interface{}) error {
		ra := compareNumber(item, a)
		rb := compareNumber(item, b)

		if ra == uncomparable || rb == uncomparable {
			return ErrInvalidType
		}

		if ra == lessThan || rb == greaterThan {
			return betweenErr
		}

		return nil
	}
}

// Required ...
func Required() Fn {
	return func(item interface{}) error {
		return nil
	}
}

// LessThan ...
func LessThan() Fn {
	return func(item interface{}) error {
		return nil
	}
}

// GreaterThan ...
func GreaterThan() Fn {
	return func(item interface{}) error {
		return nil
	}
}
