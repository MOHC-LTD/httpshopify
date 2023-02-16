package slices_test

import (
	"reflect"
	"testing"

	"github.com/MOHC-LTD/httpshopify/v2/internal/assertions"

	"github.com/MOHC-LTD/httpshopify/v2/internal/slices"
)

func TestJoinInt64(t *testing.T) {
	var input []int64 = []int64{1, 2, 3, 4}

	actual := slices.JoinInt64(input, ",")

	expected := "1,2,3,4"

	if actual != expected {
		assertions.ValueAssertionFailure(t, expected, actual)
	}
}

func TestSplitInt64(t *testing.T) {
	var input string = "1,2,3,4"

	actual := slices.SplitInt64(input, ",")

	expected := []int64{1, 2, 3, 4}

	if !reflect.DeepEqual(actual, expected) {
		assertions.ValueAssertionFailure(t, expected, actual)
	}
}
