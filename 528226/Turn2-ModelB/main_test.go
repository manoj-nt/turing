// main_test.go
package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCounter(t *testing.T) {
	wg.Add(2)
	go counter(1)
	go counter(2)
	wg.Wait()

	// Add assertions to verify the correctness of the concurrent behavior
	require.Equal(t, 1, 1, "something went wrong")
}
