package leet_code

import "unicode"

func isPalindrome(s string) bool {
	var runes []rune

	for _, r := range s {
		if unicode.IsLetter(r) {
			runes = append(runes, unicode.ToLower(r))
		} else if unicode.IsDigit(r) {
			runes = append(runes, r)
		}
	}

	if l := len(runes); l > 0 {
		for i := 0; i < l/2; i++ {
			if runes[i] != runes[l-i-1] {
				return false
			}
		}
	}
	return true

}

//func isPalindrome(s string) bool {
//	var runes []rune
//
//	for _, r := range s {
//		if unicode.IsLetter(r) {
//			runes = append(runes, unicode.ToLower(r))
//		} else if unicode.IsDigit(r) {
//			runes = append(runes, r)
//		}
//	}
//
//	if l := len(runes); l > 0 {
//		for i := 0; i < l/2; i++ {
//			if runes[i] != runes[l-i-1] {
//				return false
//			}
//		}
//	}
//	return true
//
//}
