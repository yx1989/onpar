package matchers

import (
	"fmt"
	"reflect"
)

type HaveCapMatcher struct {
	expected int
}

func HaveCap(expected int) HaveCapMatcher {
	return HaveCapMatcher{
		expected: expected,
	}
}

func (m HaveCapMatcher) Match(actual interface{}) (interface{}, error) {
	var c int
	switch reflect.TypeOf(actual).Kind() {
	case reflect.Slice, reflect.Array, reflect.Map, reflect.String, reflect.Chan:
		c = reflect.ValueOf(actual).Cap()
	default:
		return nil, fmt.Errorf("'%v' (%T) is not a Slice, Array, or Channel", actual, actual)
	}

	if c != m.expected {
		return nil, fmt.Errorf("%v (cap=%d) does not have a capacity f %d", actual, c, m.expected)
	}

	return actual, nil
}
