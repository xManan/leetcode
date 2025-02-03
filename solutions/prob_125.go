package solutions

import (
	"strings"
	"unicode"
)

// a(97) z(122) A(65) Z(90) 0(48) 9(57)
func IsPalindrome(s string) bool {
    var sb strings.Builder
    for _, r := range s {
        if unicode.IsDigit(r) {
            sb.WriteRune(r)
        } else if unicode.IsLetter(r) {
            if unicode.IsLower(r) {
                sb.WriteRune(r)
            } else {
                sb.WriteRune(unicode.ToLower(r))
            }
        }
    }
    str := sb.String()
    i := 0
    j := len(str) - 1
    for j > i {
        if str[i] != str[j] {
            return false
        }
        i++
        j--
    }
    return true
}
