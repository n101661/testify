package testify

import (
	"reflect"

	"github.com/stretchr/testify/assert"
)

func RegisterComparison(t reflect.Type, f assert.ComparisonFunc) (exist bool) {
	return assert.RegisterComparison(t, f)
}
