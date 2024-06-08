package strings

import "testing"

type testCase struct {
	description string
	input       string
	expected    string
}

func TestToPascalCase(t *testing.T) {
	inputs := []testCase{
		{
			description: "simple string",
			input:       "main",
			expected:    "Main",
		},
		{
			description: "with _",
			input:       "create_group",
			expected:    "CreateGroup",
		},
	}
	for _, tc := range inputs {
		t.Run(tc.description, func(t *testing.T) {
			result := ToPascalCase(tc.input)
			if result != tc.expected {
				t.Errorf("expected: %v got: %v", tc.expected, result)
			}
		})
	}
}
