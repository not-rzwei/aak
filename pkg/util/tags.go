package util

func HaveTags(tags1, tags2 []string) bool {
	hash := make(map[string]bool)

	for _, e := range tags1 {
		hash[e] = true
	}

	for _, e := range tags2 {
		if hash[e] {
			return true
		}
	}

	return false
}
