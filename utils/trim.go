package utils

// Trim the FIRST character in the string
func TrimF(s string) string {
	return s[1:]
}

// Trim the LAST character in the string
func TrimL(s string) string {
	return s[:len(s)-1]
}

// Trim the FIRST & LAST character in the string
func TrimFL(s string) string {
	str := s[:len(s)-1]
	return str[1:]
}
