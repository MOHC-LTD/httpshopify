package assertions

import (
	"testing"
)

// AssertionFailure provides pretty printing for a simple assertion failure
func AssertionFailure(t *testing.T, message string) {
	t.Errorf("\n%s", message)
}

// ValueAssertionFailure provides pretty printing for assertion failures
func ValueAssertionFailure(t *testing.T, expected, actual interface{}) {
	t.Errorf(
		"\nActual does not equal expected:\n\n"+
			"\tExpected: %+v\n\n"+
			"\tActual: %+v", expected, actual,
	)
}

// TypeAssertionFailure provides pretty printing for type assertion failures
func TypeAssertionFailure(t *testing.T, expected, actual interface{}) {
	t.Errorf(
		"\nActual type does not equal expected type:\n\n"+
			"\tExpected: %T\n\n"+
			"\tActual: %T", expected, actual,
	)
}

// ErrAssertionFailure provides pretty printing for unexpected error failures
func ErrAssertionFailure(t *testing.T, receivedErr error) {
	t.Errorf("\nReceived unexpected error: \n\n%v", receivedErr)
}
