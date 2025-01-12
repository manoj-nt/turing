package example

import (
	"testing"
	"time"

	"modelb/testutils"
)

func TestSomething(t *testing.T) {
	testutils.TestTimeout(t, 10*time.Second)
	// Your test code here
}

func TestSomethingElse(t *testing.T) {
	// By default, this test will use the default timeout
	testutils.TestTimeout(t, 0)
	// Your test code here
}
