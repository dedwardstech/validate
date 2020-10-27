package validate

import (
	"reflect"
	"strconv"
)

type comparision int

const (
	greaterThan comparision = iota
	lessThan
	equal
	uncomparable
)

func compareLength(item interface{}, toCompare int) comparision {
	v := reflect.ValueOf(item)
	t := reflect.TypeOf(item)

	switch t.Kind() {
	case reflect.String, reflect.Slice:
		itemLen := v.Len()

		if itemLen < toCompare {
			return lessThan
		} else if itemLen > toCompare {
			return greaterThan
		} else {
			return equal
		}
	default:
		return uncomparable
	}
}

func compareNumber(item interface{}, other int) comparision {
	v := reflect.ValueOf(item)
	t := reflect.TypeOf(item)

	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		a, b := v.Int(), int64(other)

		if a < b {
			return lessThan
		} else if a > b {
			return greaterThan
		} else {
			return equal
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		a, b := v.Uint(), uint64(other)
		if a < b {
			return lessThan
		} else if a > b {
			return greaterThan
		} else {
			return equal
		}
	case reflect.Float32, reflect.Float64:
		a, b := v.Float(), float64(other)
		if a < b {
			return lessThan
		} else if a > b {
			return greaterThan
		} else {
			return equal
		}
	case reflect.String:
		a, b := v.String(), float64(other)
		val, err := strconv.ParseFloat(a, 64)
		if err != nil {
			return uncomparable
		}

		if val < b {
			return lessThan
		} else if val > b {
			return greaterThan
		} else {
			return equal
		}
	default:
		return uncomparable
	}
}
