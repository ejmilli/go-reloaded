package main

func removeTheSlice(remove []string) []string {

	var result []string

	for _, char := range remove {
		if char != "" {
			result = append(result, char)
		}

	}
	return result
}