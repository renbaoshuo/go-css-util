package cssutil

// This function checks if two code points can start a valid escape
// sequence in CSS.
//
//   - If the first code point is not U+005C REVERSE SOLIDUS (\),
//     return false.
//   - Otherwise, if the second code point is a newline, return false.
//   - Otherwise, return true.
//
// https://www.w3.org/TR/css-syntax-3/#check-if-two-code-points-are-a-valid-escape
func TwoCodePointsStartsAValidEscape(first, second rune) bool {
	return first == '\\' && !IsNewline(second)
}
