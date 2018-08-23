package test

import (
	"testing"

	"github.com/tecnologer/romantoarabic/num"
)

type romanTest struct {
	input  string
	output int
}

func TestRoman(t *testing.T) {
	var romant = []romanTest{
		romanTest{
			input:  "i",
			output: 1,
		},
		romanTest{
			input:  "MMMDCCCLXXXVIII",
			output: 3888,
		},
		romanTest{
			input:  "v",
			output: 5,
		},
		romanTest{
			input:  "V",
			output: 5,
		},
		romanTest{
			input:  "a",
			output: -1,
		},
	}

	for _, v := range romant {
		output := -1
		if num.IsRoman(v.input) {
			output = num.ParseRoman(v.input)
		}

		if v.output != output {
			t.Errorf("Error al convertir %s", v.input)
		}
	}
}
