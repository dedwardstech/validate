package validate

import (
	"reflect"
	"testing"

	"github.com/pkg/errors"
)

func Test_MinLength(t *testing.T) {
	wantedErr := errors.New("must be longer than 6")
	cases := []struct {
		arg interface{}
		err error
	}{
		{arg: []int{1, 2, 3, 4, 5}, err: wantedErr},
		{arg: "abcde", err: wantedErr},
		{arg: []int{1, 2, 4, 4, 5, 6}, err: nil},
		{arg: []int{1, 2, 4, 4, 5, 6, 7}, err: nil},
		{arg: "abcdef", err: nil},
		{arg: "abcdefg", err: nil},
	}

	validate := MinLength(6, "must be longer than 6")
	for _, c := range cases {
		err := validate(c.arg)

		if c.err == nil && err != nil {
			t.Errorf("should not return an error for %v", c.arg)
		} else if c.err != nil && err == nil {
			t.Errorf("wanted err: %v, for %v, type=%v", c.err, c.arg, reflect.TypeOf(c.arg))
		} else if c.err != nil && err != nil && c.err.Error() != err.Error() {
			t.Errorf("wanted error: %v, got error: %v", c.err.Error(), err.Error())
		}
	}
}

func Test_MaxLength(t *testing.T) {
	wantedErr := errors.New("must be less than 4")
	cases := []struct {
		arg interface{}
		err error
	}{
		{arg: []int{1, 2, 3, 4, 5}, err: wantedErr},
		{arg: []string{"1", "2", "3", "4", "5"}, err: wantedErr},
		{arg: []int{1, 2, 3, 4}, err: nil},
		{arg: []int{1, 2, 3}, err: nil},
	}

	validate := MaxLength(4, "must be less than 4")
	for _, c := range cases {
		err := validate(c.arg)

		if c.err == nil && err != nil {
			t.Errorf("MaxLength -> should not return an error")
		} else if c.err != nil && err == nil {
			t.Errorf("MaxLength -> wanted err: %v, for %v, type=%v", c.err, c.arg, reflect.TypeOf(c.arg))
		} else if c.err != nil && err != nil && c.err.Error() != err.Error() {
			t.Errorf("MaxLength -> wanted error: %v, got error: %v", c.err.Error(), err.Error())
		}
	}
}

func Test_Between(t *testing.T) {
	wantedErr := errors.New("must be between 3 and 5")
	cases := []struct {
		arg interface{}
		err error
	}{
		{arg: 2, err: wantedErr},
		{arg: uint(2), err: wantedErr},
		{arg: uint8(2), err: wantedErr},
		{arg: uint16(2), err: wantedErr},
		{arg: uint32(2), err: wantedErr},
		{arg: uint64(2), err: wantedErr},
		{arg: float32(2.99999), err: wantedErr},
		{arg: float64(2.9999999), err: wantedErr},
		{arg: "2", err: wantedErr},
		{arg: "2.99999999", err: wantedErr},
		{arg: 6, err: wantedErr},
		{arg: uint(6), err: wantedErr},
		{arg: uint8(6), err: wantedErr},
		{arg: uint16(6), err: wantedErr},
		{arg: uint32(6), err: wantedErr},
		{arg: uint64(6), err: wantedErr},
		{arg: float32(5.99999), err: wantedErr},
		{arg: float64(5.9999999), err: wantedErr},
		{arg: "6", err: wantedErr},
		{arg: "5.99999999", err: wantedErr},
		{arg: 3, err: nil},
		{arg: uint(3), err: nil},
		{arg: uint8(3), err: nil},
		{arg: uint16(3), err: nil},
		{arg: uint32(3), err: nil},
		{arg: uint64(3), err: nil},
		{arg: float32(3.00001), err: nil},
		{arg: float64(3.00000001), err: nil},
		{arg: "3", err: nil},
		{arg: "3.0000000001", err: nil},
		{arg: 4, err: nil},
		{arg: uint(4), err: nil},
		{arg: uint8(4), err: nil},
		{arg: uint16(4), err: nil},
		{arg: uint32(4), err: nil},
		{arg: uint64(4), err: nil},
		{arg: float32(4.00001), err: nil},
		{arg: float64(4.00000001), err: nil},
		{arg: "4", err: nil},
		{arg: "4.0000000001", err: nil},
		{arg: 5, err: nil},
		{arg: uint(5), err: nil},
		{arg: uint8(5), err: nil},
		{arg: uint16(5), err: nil},
		{arg: uint32(5), err: nil},
		{arg: uint64(5), err: nil},
		{arg: float32(4.99991), err: nil},
		{arg: float64(4.99999991), err: nil},
		{arg: "5", err: nil},
		{arg: "4.99999999991", err: nil},
	}

	validate := Between(3, 5, "must be between 3 and 5")
	for _, c := range cases {
		err := validate(c.arg)

		if c.err == nil && err != nil {
			t.Errorf("should not return an error")
		} else if c.err != nil && err == nil {
			t.Errorf("wanted err: %v, for %v, type=%v", c.err, c.arg, reflect.TypeOf(c.arg))
		} else if c.err != nil && err != nil && c.err.Error() != err.Error() {
			t.Errorf("wanted error: %v, got error: %v", c.err.Error(), err.Error())
		}
	}
}

func Test_Required(t *testing.T) {

}

func Test_LessThan(t *testing.T) {

}

func Test_GreaterThan(t *testing.T) {

}
