package helpers

import (
	"unicode"
	"unicode/utf8"
)

func FirstToUpper(s string) string {
	r, size := utf8.DecodeRuneInString(s)
	if r == utf8.RuneError && size <= 1 {
		return s
	}
	fc := unicode.ToUpper(r)
	if r == fc {
		return s
	}
	return string(fc) + s[size:]
}
