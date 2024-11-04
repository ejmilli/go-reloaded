package main

func toUpper(up string) string {
	var result string
	for _, char := range up {
		if char >= 'a' && char <= 'z' {
			char -= 32
		}
		result += string(char)
	}
	return result
}