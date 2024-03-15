package main

import "testing"

func TestCleanInput(t *testing.T){
	cases := []struct {
		input string
		expected []string
	}{
		{
			input : "hello world",
			expected : []string{
				"hello",
				"world",
			},
		},
		{
			input : "HeLlo WOrld",
			expected : []string{
				"hello",
				"world",
			},
		},
	}
	
	for _, cs := range cases {
		actual := cleanInput(cs.input)
		if len(actual) != len(cs.expected){
			t.Errorf("The lengths are not equal\n - Actual: %v\n - Expected: %v", 
			len(actual), 
			len(cs.expected),
			)
			continue
		}
		for i := range actual {
			actualWord := actual[i]
			expectedWord := cs.expected[i]
			if expectedWord != actualWord {
				t.Errorf("Expected: %v, does not equal Actual: %v", 
				expectedWord, 
				actualWord)
			}
		}
	}
}