package utils

// trim the FIRST character in the string
func TrimF(s string) string { return s[1:len(s)] }

// trim the LAST character in the string
func TrimL(s string) string { return s[:len(s)-1] }

// trim the FIRST & LAST character in the string
func TrimFL(s string) string {
	str := s[:len(s)-1]
	return str[1:len(str)]
}
