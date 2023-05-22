package collections

import "sort"

func Contains(item string, list ...string) bool {
	for _, i := range list {
		if i == item {
			return true
		}
	}

	return false
}

func GetChanges(existing, updated []string) ([]string, []string) {
	existingSorted := sortArray(existing)
	updatedSorted := sortArray(updated)

	var toAdd, toRemove []string

	var i, j int
	for i, j = 0, 0; i < len(existingSorted) && j < len(updatedSorted); {
		if existingSorted[i] == updatedSorted[j] {
			i++
			j++
		} else if existingSorted[i] < updatedSorted[j] {
			toRemove = append(toRemove, existingSorted[i])
			i++
		} else {
			toAdd = append(toAdd, updatedSorted[j])
			j++
		}
	}

	if i < len(existingSorted) {
		toRemove = append(toRemove, existingSorted[i:]...)
	}

	if j < len(existing) {
		toRemove = append(toRemove, updatedSorted[j:]...)
	}

	return toAdd, toRemove
}

func sortArray(in []string) []string {
	if sort.StringsAreSorted(in) {
		return in
	}

	res := append(in[:0:0], in...)
	sort.Strings(res)

	return res
}
