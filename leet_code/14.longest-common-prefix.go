package leet_code

// https://leetcode.com/problems/longest-common-prefix

func longestCommonPrefix(strs []string) string {
	var prefix []byte
	for i := 0; i < len(strs); i++ {
		if i == 0 {
			prefix = []byte(strs[i])
			continue
		}
		if len(prefix) == 0 {
			break
		}
		word := []byte(strs[i])
		for j := 0; j < len(word); j++ {
			if j >= len(prefix) {
				break
			}
			if prefix[j] != word[j] {
				prefix = prefix[:j]
			}
		}
		if len(word) < len(prefix) {
			prefix = prefix[:len(word)]
		}
	}

	return string(prefix)
}
