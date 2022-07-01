package reflection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type reflectionTestCase struct {
	message        string
	input          interface{}
	expectedFields []string
}

type StructWithStringAndInt struct {
	A string
	B int
}

type NestedStruct struct {
	A string
	B StructWithStringAndInt
}

func TestWalk(t *testing.T) {
	testCases := []reflectionTestCase{
		{
			message: "two fields flat form",
			input: struct {
				a string
				b string
			}{"123", "456"},
			expectedFields: []string{"123", "456"},
		},
		{
			message: "three fields flat form",
			input: struct {
				a string
				b string
				c string
			}{"123", "456", "234"},
			expectedFields: []string{"123", "456", "234"},
		},
		{
			message: "with number",
			input: StructWithStringAndInt{
				A: "123",
				B: 456,
			},
			expectedFields: []string{"123"},
		},
		{
			message: "nested",
			input: NestedStruct{
				A: "123",
				B: StructWithStringAndInt{
					A: "456",
					B: 1234,
				},
			},
			expectedFields: []string{"123", "456"},
		},
		{
			message:        "pointer to struct",
			input:          &struct{ a string }{"123"},
			expectedFields: []string{"123"},
		},
		{
			message: "slice",
			input: []StructWithStringAndInt{{
				A: "123",
				B: 456,
			}, {
				A: "789",
				B: 123,
			}},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.message, func(t *testing.T) {
			var actual []string
			Walk(testCase.input, func(s string) {
				actual = append(actual, s)
			})
			assert.Equal(t, testCase.expectedFields, actual)
		})

	}

}
