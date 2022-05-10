package utils

func TrimFL(s string) string {
	str := s[:len(s)-1]

	return str[1:len(str)]
}
