package latinize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	tests := map[string]struct {
		input string

		expected string
	}{
		"Empty input": {
			"",
			"",
		},

		"Special characters": {
			"ỆᶍǍᶆṔƚÉ áéíóúýčďěňřšťžů",
			"ExAmPlE aeiouycdenrstzu",
		},

		"German umlauts #1": {
			"das Mädchen, schön, küssen",
			"das Madchen, schon, kussen",
		},

		"German umlauts #2": {
			"Yoùr Śtßring",
			"Your Stssring",
		},

		"Diacritics #1": {
			"âîûäöüąęś źżółø ıćńóśźżąęł",
			"aiuaouaes zzolo icnoszzael",
		},

		"Diacritics #2": {
			"ëïé èöêüçàû îñäôíú",
			"eie eoeucau inaoiu",
		},

		"Diacritics #3": {
			"Yoùr Śtring šđčćžŠĐČĆŽ Ötzi's Nationalität èàì",
			"Your String sdcczSDCCZ Otzi's Nationalitat eai",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := String(test.input)

			assert.Equal(t, test.expected, actual)
		})
	}
}
