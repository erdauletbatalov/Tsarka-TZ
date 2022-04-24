package usecase

import (
	"testing"
)

func Test_FindSubstr(t *testing.T) {
	testCases := []struct {
		name     string
		in       string
		expected string
	}{
		{
			name:     "one",
			in:       "abobaaboba",
			expected: "abo",
		},
		{
			name:     "two",
			in:       "kilmelo",
			expected: "kilme",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if LongestSubstring(tc.in) != tc.expected {
				t.Errorf("LongestSubstring -> expected: %s, actual: %s", tc.expected, LongestSubstring(tc.in))
			}
		})
	}
}
