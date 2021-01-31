package internal

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func TrimRightSpace(s string) string {
	return strings.TrimRightFunc(s, unicode.IsSpace)
}

func RemoveBOM(bytes []byte) []byte {
	if len(bytes) >= 3 {
		if bytes[0] == 0xef && bytes[1] == 0xbb && bytes[2] == 0xbf {
			return bytes[3:]
		}
	}
	return bytes
}

func Split(s, sep string) []string {
	return genSplit(s, sep, 0, -1)
}

func genSplit(s, sep string, sepSave, n int) []string {
	if n == 0 {
		return nil
	}
	if sep == "" {
		return explode(s, n)
	}
	if n < 0 {
		n = strings.Count(s, sep) + 1
	}

	a := make([]string, 0, n)
	n--
	i := 0
	for i < n {
		m := strings.Index(s, sep)
		if m < 0 {
			break
		}
		var ns = s[:m+sepSave]
		s = s[m+len(sep):]
		i++

		if ns != "" {
			a = append(a, ns)
		}
	}
	a = append(a, s)
	return a
}

func explode(s string, n int) []string {
	l := utf8.RuneCountInString(s)
	if n < 0 || n > l {
		n = l
	}
	a := make([]string, n)
	for i := 0; i < n-1; i++ {
		ch, size := utf8.DecodeRuneInString(s)
		a[i] = s[:size]
		s = s[size:]
		if ch == utf8.RuneError {
			a[i] = string(utf8.RuneError)
		}
	}
	if n > 0 {
		a[n-1] = s
	}
	return a
}
