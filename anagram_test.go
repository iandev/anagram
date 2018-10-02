package main

import (
	"fmt"
	"testing"
)

type testConfig struct {
	str1     string
	str2     string
	expected bool
}

var tests = []testConfig{
	testConfig{
		str1:     "iceman",
		str2:     "cinema",
		expected: true,
	},
	testConfig{
		str1:     "tar",
		str2:     "rat",
		expected: true,
	},
	testConfig{
		str1:     "arc",
		str2:     "car",
		expected: true,
	},
	testConfig{
		str1:     "elbow",
		str2:     "below",
		expected: true,
	},
	testConfig{
		str1:     "state",
		str2:     "taste",
		expected: true,
	},
	testConfig{
		str1:     "state",
		str2:     "taste",
		expected: true,
	},
	testConfig{
		str1:     "BRag",
		str2:     "graB",
		expected: true,
	},
	testConfig{
		str1:     "nnight",
		str2:     "thing",
		expected: false,
	},
	testConfig{
		str1:     "Bored",
		str2:     "Robedd",
		expected: false,
	},
	testConfig{
		str1:     "desserts ",
		str2:     " stressed",
		expected: true,
	},
	testConfig{
		str1:     "aaaaaa",
		str2:     "abcdef",
		expected: false,
	},
}

func TestAnagram(t *testing.T) {
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			if anagram(tt.str1, tt.str2) != tt.expected {
				t.Errorf("anagram failed for %s and %s, expected %t", tt.str1, tt.str2, tt.expected)
			}
		})
	}
}
