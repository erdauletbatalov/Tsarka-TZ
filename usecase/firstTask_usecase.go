package usecase

func LongestSubstring(str string) string {
	res := ""

	for i := 0; i < len(str); i++ {
		for j := i; j < len(str); j++ {
			if IsUnique(str[i : j+1]) {
				if len(str[i:j+1]) > len(res) {
					res = str[i : j+1]
				}
			}
		}
	}
	return res
}

func IsUnique(str string) bool {
	for i := 0; i < len(str); i++ {
		for j := 0; j < len(str); j++ {
			if i != j {
				if str[i] == str[j] {
					return false
				}
			}
		}
	}

	return true
}
