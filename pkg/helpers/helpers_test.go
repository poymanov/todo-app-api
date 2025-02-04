package helpers_test

import (
	"github.com/stretchr/testify/require"
	"poymanov/todo/pkg/helpers"
	"testing"
)

func TestFirstToUpperSuccess(t *testing.T) {
	testCases := []struct {
		value    string
		expected string
	}{
		{"test", "Test"},
		{"тест", "Тест"},
		{"error description", "Error description"},
		{"1", "1"},
	}

	for _, tc := range testCases {
		t.Run("case", func(t *testing.T) {
			result := helpers.FirstToUpper(tc.value)
			require.Equal(t, tc.expected, result)
		})
	}
}
