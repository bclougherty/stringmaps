package stringmaps

func Keys(m map[string]string) []string {
	keys := make([]string, 0, len(m))

	for key, _ := range m {
		keys = append(keys, key)
	}

	return keys
}

func Values(m map[string]string) []string {
	values := make([]string, 0, len(m))

	for _, value := range m {
		values = append(values, value)
	}

	return values
}

func Reverse(m map[string]string) map[string]string {
	m2 := make(map[string]string, len(m))

	for key, value := range m {
		m2[value] = key
	}

	return m2
}
