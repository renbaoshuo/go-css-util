package cssutil

import (
	"fmt"
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
