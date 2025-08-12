package cssutil

import (
	"fmt"
	"strings"
)

// To escape a character means to create a string of "\" (U+005C), followed
// by the character.
//
// https://www.w3.org/TR/2021/WD-cssom-1-20210826/#escape-a-character
func EscapeCharacter(c rune) string {
	return "\\" + string(c)
}

// To escape a character as code point means to create a string of "\"
// (U+005C), followed by the Unicode code point as the smallest possible
// number of hexadecimal digits in the range 0-9 a-f (U+0030 to U+0039
// and U+0061 to U+0066) to represent the code point in base 16, followed
// by a single SPACE (U+0020).
//
// https://www.w3.org/TR/2021/WD-cssom-1-20210826/#escape-a-character-as-code-point
func EscapeCharacterAsCodePoint(c rune) string {
	return fmt.Sprintf("\\%x ", c)
}

// To serialize an identifier means to create a string represented by the
// concatenation of, for each character of the identifier:
//
//   - If the character is NULL (U+0000), then the REPLACEMENT CHARACTER
//     (U+FFFD).
//   - If the character is in the range [\1-\1f] (U+0001 to U+001F) or is
//     U+007F, then the character escaped as code point.
//   - If the character is the first character and is in the range [0-9]
//     (U+0030 to U+0039), then the character escaped as code point.
//   - If the character is the second character and is in the range [0-9]
//     (U+0030 to U+0039) and the first character is a "-" (U+002D), then
//     the character escaped as code point.
//   - If the character is the first character and is a "-" (U+002D), and
//     there is no second character, then the escaped character.
//   - If the character is not handled by one of the above rules and is
//     greater than or equal to U+0080, is "-" (U+002D) or "_" (U+005F),
//     or is in one of the ranges [0-9] (U+0030 to U+0039), [A-Z] (U+0041
//     to U+005A), or [a-z] (U+0061 to U+007A), then the character itself.
//   - Otherwise, the escaped character.
//
// https://www.w3.org/TR/2021/WD-cssom-1-20210826/#serialize-an-identifier
func SerializeIdentifier(identifier string) string {
	var result strings.Builder

	runes := []rune(identifier)

	for i, c := range runes {
		switch {
		case c == 0x0000: // NULL character
			result.WriteRune(0xFFFD) // REPLACEMENT CHARACTER
		case (c >= 0x0001 && c <= 0x001F) || c == 0x007F: // Control characters
			result.WriteString(EscapeCharacterAsCodePoint(c))
		case i == 0 && IsDigit(c): // First character is a digit
			result.WriteString(EscapeCharacterAsCodePoint(c))
		case i == 1 && IsDigit(c) && runes[0] == '-': // Second character is digit and first is hyphen
			result.WriteString(EscapeCharacterAsCodePoint(c))
		case i == 0 && c == '-' && len(runes) == 1: // First and only character is hyphen
			result.WriteString(EscapeCharacter(c))
		case IsNonASCII(c) || // Non-ASCII
			c == '-' || // Hyphen (-)
			c == '_' || // Underscore (_)
			IsDigit(c) || // Digits (0-9)
			IsLetter(c): // Letters (A-Z, a-z)
			result.WriteRune(c)
		default: // All other characters need to be escaped
			result.WriteString(EscapeCharacter(c))
		}
	}

	return result.String()
}

// To serialize a string means to create a string represented by '"'
// (U+0022), followed by the result of applying the rules below to each
// character of the given string, followed by '"' (U+0022):
//
//   - If the character is NULL (U+0000), then the REPLACEMENT CHARACTER
//     (U+FFFD).
//   - If the character is in the range [\1-\1f] (U+0001 to U+001F) or is
//     U+007F, the character escaped as code point.
//   - If the character is '"' (U+0022) or "\" (U+005C), the escaped
//     character.
//   - Otherwise, the character itself.
//
// Note: "'" (U+0027) is not escaped because strings are always serialized
// with '"' (U+0022).
//
// https://www.w3.org/TR/2021/WD-cssom-1-20210826/#serialize-a-string
func SerializeString(s string) string {
	var result strings.Builder

	result.WriteRune('"') // Start with opening quote

	for _, c := range s {
		switch {
		case c == 0x0000: // NULL character
			result.WriteRune(0xFFFD) // REPLACEMENT CHARACTER
		case (c >= 0x0001 && c <= 0x001F) || c == 0x007F: // Control characters
			result.WriteString(EscapeCharacterAsCodePoint(c))
		case c == '"' || c == '\\': // Quote or backslash
			result.WriteString(EscapeCharacter(c))
		default:
			result.WriteRune(c)
		}
	}

	result.WriteRune('"') // End with closing quote

	return result.String()
}

// To serialize a URL means to create a string represented by "url(",
// followed by the serialization of the URL as a string, followed by ")".
//
// https://www.w3.org/TR/2021/WD-cssom-1-20210826/#serialize-a-url
func SerializeURL(url string) string {
	return fmt.Sprintf("url(%s)", SerializeString(url))
}

// To serialize a LOCAL means to create a string represented by "local(",
// followed by the serialization of the LOCAL as a string, followed by ")".
//
// https://www.w3.org/TR/2021/WD-cssom-1-20210826/#serialize-a-local
func SerializeLocal(local string) string {
	return fmt.Sprintf("local(%s)", SerializeString(local))
}

// To serialize a comma-separated list concatenate all items of the list in
// list order while separating them by ", ", i.e., COMMA (U+002C) followed
// by a single SPACE (U+0020).
//
// https://www.w3.org/TR/2021/WD-cssom-1-20210826/#serialize-a-comma-separated-list
func SerializeCommaSeparatedList(items []string) string {
	return strings.Join(items, ", ")
}

// To serialize a whitespace-separated list concatenate all items of the
// list in list order while separating them by " ", i.e., a single SPACE
// (U+0020).
//
// https://www.w3.org/TR/2021/WD-cssom-1-20210826/#serialize-a-whitespace-separated-list
func SerializeWhitespaceSeparatedList(items []string) string {
	return strings.Join(items, " ")
}
