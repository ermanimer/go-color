package color

import (
	"testing"
)

func TestMakeEscCode(t *testing.T) {
	tests := []struct {
		name      string
		typeCode  typeCode
		colorCode colorCode
		expected  escCode
	}{
		{
			name:      "fg",
			typeCode:  fg,
			colorCode: White,
			expected:  "\x1b[38;5;7;m",
		},
		{
			name:      "bg",
			typeCode:  bg,
			colorCode: Red,
			expected:  "\x1b[48;5;1;m",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			actual := makeEscCode(test.typeCode, test.colorCode)
			if actual != test.expected {
				tt.Errorf("%s != %s", actual, test.expected)
			}
		})
	}
}
