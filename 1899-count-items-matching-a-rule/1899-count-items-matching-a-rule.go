func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	hashMap := make(map[string]int)
	hashMap["type"] = 0
	hashMap["color"] = 1
	hashMap["name"] = 2

	var res int
	for _, item := range items {
		if item[hashMap[ruleKey]] == ruleValue {
			res++
		}
	}
	return res
}