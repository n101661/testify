package assert

import (
	"reflect"
)

type EqualComparisonFunc func(expected, actual interface{}) bool

var registeredEqualComparisons = map[reflect.Type]EqualComparisonFunc{}

func RegisterEqualComparison(t reflect.Type, f EqualComparisonFunc) (exist bool) {
	if _, exist = registeredEqualComparisons[t]; exist {
		return
	}
	registeredEqualComparisons[t] = f
	return false
}
