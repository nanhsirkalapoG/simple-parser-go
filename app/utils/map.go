package utils

// Check any one of the array elements is present in a hash map and return the index(-1 for no match)
func GetMatchedIndex(elements []string, fieldMap map[string]int) int {
	for _, value := range elements {
		if fieldMap[value] != 0 {
			return fieldMap[value] - 1
		}
	}
	return -1
}
