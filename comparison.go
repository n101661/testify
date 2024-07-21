package testify

import (
	"reflect"

	"github.com/stretchr/testify/assert"
)

func RegisterEqualComparison(t reflect.Type, f assert.EqualComparisonFunc) (exist bool) {
	return assert.RegisterEqualComparison(t, f)
}
