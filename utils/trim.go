package utils

func TrimF(s string) string { return s[1:len(s) }

func TrimL(s string) string { return s[:len(s)-1] }

func TrimFL(s string) string {
	str := s[:len(s)-1]

	return str[1:len(str)]
}
