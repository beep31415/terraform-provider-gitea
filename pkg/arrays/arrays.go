package arrays

func Contains(item string, list ...string) bool {
	for _, i := range list {
		if i == item {
			return true
		}
	}

	return false
}
