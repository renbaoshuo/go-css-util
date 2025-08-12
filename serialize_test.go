package cssutil

import (
	"testing"
)

func TestEscapeCharacter(t *testing.T) {
	tests := []struct {
		input    rune
		expected string
	}{
		{'a', `\a`},
		{'A', `\A`},
		{'0', `\0`},
		{'_', `\_`},
		{' ', `\ `},
	}

	for _, test := range tests {
		result := EscapeCharacter(test.input)
		if result != test.expected {
			t.Errorf("EscapeCharacter(%c) = %s; want %s", test.input, result, test.expected)
		}
	}
}

func TestEscapeCharacterAsCodePoint(t *testing.T) {
	tests := []struct {
		input    rune
		expected string
	}{
		{'a', `\61 `},
		{'A', `\41 `},
		{'0', `\30 `},
		{'_', `\5f `},
		{' ', `\20 `},
		{'$', `\24 `},
		{'©', `\a9 `},    // U+00A9 COPYRIGHT SIGN
		{'€', `\20ac `},  // U+20AC EURO SIGN
		{'𐍈', `\10348 `}, // U+10348
	}

	for _, test := range tests {
		result := EscapeCharacterAsCodePoint(test.input)
		if result != test.expected {
			t.Errorf("EscapeCharacterAsCodePoint(%c) = %s; want %s", test.input, result, test.expected)
		}
	}
}
