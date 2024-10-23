package main

func ToLower(low string) string {
	result := ""
	for _, char := range low {
		if char >= 'a' && char <= 'z' {
			char += 32
		}
		result += string(char)
	}
	return result
}