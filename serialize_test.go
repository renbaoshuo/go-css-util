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

func TestSerializeString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// Basic string
		{"hello", `"hello"`},
		// Empty string
		{"", `""`},
		// String with quotes
		{`hello "world"`, `"hello \"world\""`},
		// String with backslash
		{`hello\world`, `"hello\\world"`},
		// String with NULL character
		{string(rune(0x0000)), `"�"`}, // REPLACEMENT CHARACTER
		// String with control characters
		{string(rune(0x0001)), `"\1 "`},
		{string(rune(0x001F)), `"\1f "`},
		{string(rune(0x007F)), `"\7f "`},
		// Mixed content
		{`test"123\control` + string(rune(0x0001)), `"test\"123\\control\1 "`},
	}

	for _, test := range tests {
		result := SerializeString(test.input)
		if result != test.expected {
			t.Errorf("SerializeString(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestSerializeIdentifier(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// Basic identifier
		{"hello", "hello"},
		{"Hello", "Hello"},
		{"_test", "_test"},
		{"test123", "test123"},
		{"test-case", "test-case"},
		// Identifier starting with digit (should be escaped)
		{"1test", `\31 test`},
		{"0abc", `\30 abc`},
		// Identifier starting with hyphen followed by digit
		{"-1test", `-\31 test`},
		{"-0abc", `-\30 abc`},
		// Single hyphen (should be escaped)
		{"-", `\-`},
		// Hyphen followed by letter (not escaped)
		{"-test", "-test"},
		// NULL character
		{string(rune(0x0000)), "�"}, // REPLACEMENT CHARACTER
		// Control characters
		{string(rune(0x0001)), `\1 `},
		{string(rune(0x001F)), `\1f `},
		{string(rune(0x007F)), `\7f `},
		// Non-ASCII characters (should not be escaped)
		{"café", "café"},
		{"测试", "测试"},
		// Special characters that need escaping
		{"test@example", `test\@example`},
		{"test()", `test\(\)`},
		{"test{}", `test\{\}`},
	}

	for _, test := range tests {
		result := SerializeIdentifier(test.input)
		if result != test.expected {
			t.Errorf("SerializeIdentifier(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestSerializeURL(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"http://example.com", `url("http://example.com")`},
		{"https://example.com/path", `url("https://example.com/path")`},
		{"", `url("")`},
	}

	for _, test := range tests {
		result := SerializeURL(test.input)
		if result != test.expected {
			t.Errorf("SerializeURL(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestSerializeLocal(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Arial", `local("Arial")`},
		{"Times New Roman", `local("Times New Roman")`},
		{"", `local("")`},
	}

	for _, test := range tests {
		result := SerializeLocal(test.input)
		if result != test.expected {
			t.Errorf("SerializeLocal(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}
