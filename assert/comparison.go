package assert

import (
	"reflect"
)

type ComparisonFunc func(e1, e2 interface{}) int

var registeredComparisons = map[reflect.Type]ComparisonFunc{}

func RegisterComparison(t reflect.Type, f ComparisonFunc) (exist bool) {
	if _, exist = registeredComparisons[t]; exist {
		return
	}
	registeredComparisons[t] = f
	return false
}
