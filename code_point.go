package cssutil

// This file contains definitions of code points used in CSS syntax.
// These utilities are useful for tokenization and parsing of CSS.
//
// https://www.w3.org/TR/2021/CRD-css-syntax-3-20211224/#tokenizer-definitions

// A code point between U+0030 DIGIT ZERO (0) and U+0039 DIGIT NINE (9)
// inclusive.
//
// https://www.w3.org/TR/2021/CRD-css-syntax-3-20211224/#digit
func IsDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

// A digit [IsDigit], or a code point between U+0041 LATIN CAPITAL LETTER
// A (A) and U+0046 LATIN CAPITAL LETTER F (F) inclusive, or a code point
// between U+0061 LATIN SMALL LETTER A (a) and U+0066 LATIN SMALL LETTER
// F (f) inclusive.
//
// https://www.w3.org/TR/2021/CRD-css-syntax-3-20211224/#hex-digit
func IsHexDigit(c rune) bool {
	return IsDigit(c) || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')
}

// A code point between U+0041 LATIN CAPITAL LETTER A (A) and U+005A LATIN
// CAPITAL LETTER Z (Z) inclusive.
//
// https://www.w3.org/TR/2021/CRD-css-syntax-3-20211224/#uppercase-letter
func IsUpperCaseLetter(c rune) bool {
	return c >= 'A' && c <= 'Z'
}

// A code point between U+0061 LATIN SMALL LETTER A (a) and U+007A LATIN
// SMALL LETTER Z (z) inclusive.
//
// https://www.w3.org/TR/2021/CRD-css-syntax-3-20211224/#lowercase-letter
func IsLowerCaseLetter(c rune) bool {
	return c >= 'a' && c <= 'z'
}

// An uppercase letter [IsUpperCaseLetter] or a lowercase letter [IsLowerCaseLetter].
//
// https://www.w3.org/TR/2021/CRD-css-syntax-3-20211224/#letter
func IsLetter(c rune) bool {
	return IsUpperCaseLetter(c) || IsLowerCaseLetter(c)
}

// A code point with a value equal to or greater than U+0080 <control>.
//
// https://www.w3.org/TR/2021/CRD-css-syntax-3-20211224/#non-ascii-code-point
func IsNonASCII(c rune) bool {
	return c >= 0x80
}

// A letter [IsLetter], a non-ASCII code point [IsNonASCII], or U+005F LOW LINE (_).
//
// https://www.w3.org/TR/2021/CRD-css-syntax-3-20211224/#ident-start-code-point
func IsIdentStartCodePoint(r rune) bool {
	return IsLetter(r) || IsNonASCII(r) || r == '_'
}

// An ident-start code point [IsIdentStartCodePoint], a digit [IsDigit], or
// U+002D HYPHEN-MINUS (-).
//
// https://www.w3.org/TR/2021/CRD-css-syntax-3-20211224/#ident-code-point
func IsIdentCodePoint(r rune) bool {
	return IsIdentStartCodePoint(r) || IsDigit(r) || r == '-'
}

// A code point between U+0000 NULL and U+0008 BACKSPACE inclusive, or
// U+000B LINE TABULATION, or a code point between U+000E SHIFT OUT and
// U+001F INFORMATION SEPARATOR ONE inclusive, or U+007F DELETE.
//
// https://www.w3.org/TR/2021/CRD-css-syntax-3-20211224/#non-printable-code-point
func IsNonPrintableCodePoint(c rune) bool {
	return (c >= 0x00 && c <= 0x08) || c == 0x0B || (c >= 0x0E && c <= 0x1F) || c == 0x7F
}

// U+000A LINE FEED.
//
// Note that U+000D CARRIAGE RETURN and U+000C FORM FEED are not included
// in this definition, as they are converted to U+000A LINE FEED during
// preprocessing.
//
// https://www.w3.org/TR/2021/CRD-css-syntax-3-20211224/#newline
func IsNewline(c rune) bool {
	return c == '\n' || c == '\r' || c == '\f'
}

// A newline [IsNewline], U+0009 CHARACTER TABULATION, or U+0020 SPACE.
//
// https://www.w3.org/TR/2021/CRD-css-syntax-3-20211224/#whitespace
func IsWhitespace(c rune) bool {
	return IsNewline(c) || c == '\t' || c == ' '
}

// The greatest code point defined by Unicode: U+10FFFF.
//
// https://www.w3.org/TR/2021/CRD-css-syntax-3-20211224/#maximum-allowed-code-point
func IsLowerThanMaxCodePoint(c rune) bool {
	return c <= 0x10FFFF
}

// A leading surrogate is a code point that is in the range U+D800 to
// U+DBFF, inclusive.
//
// https://infra.spec.whatwg.org/#leading-surrogate
func IsLeadingSurrogate(c rune) bool {
	return c >= 0xD800 && c <= 0xDBFF
}

// A trailing surrogate is a code point that is in the range U+DC00 to
// U+DFFF, inclusive.
//
// https://infra.spec.whatwg.org/#trailing-surrogate
func IsTrailingSurrogate(c rune) bool {
	return c >= 0xDC00 && c <= 0xDFFF
}

// A surrogate is a leading surrogate [IsLeadingSurrogate] or a trailing
// surrogate [IsTrailingSurrogate].
//
// https://infra.spec.whatwg.org/#surrogate
func IsSurrogate(c rune) bool {
	return IsLeadingSurrogate(c) || IsTrailingSurrogate(c)
}
