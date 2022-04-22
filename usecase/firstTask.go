package usecase

func LongestSubstring(str string) string {
	// making a map the go way
	hashLen := make(map[int]int)

	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			if j == len(str) {
				break
			}
			if contains(str[i:j], rune(str[j])) {
				if _, ok := hashLen[len(str[i:j])]; !ok {
					hashLen[len(str[i:j])] = i
				}
				break
			}
		}
	}

	keys := make([]int, 0, len(hashLen))
	for k := range hashLen {
		keys = append(keys, k)
	}
	return str[hashLen[max(keys)] : hashLen[max(keys)]+max(keys)]
}

func max(array []int) int {
	var max int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
	}
	return max
}

func contains(str1 string, char rune) bool {
	for _, val := range str1 {
		if val == char {
			return true
		}
	}
	return false
}

// for l, r := 0, 0; r < len(str); r++ {
// 	if index, ok := hashLen[str[r]]; ok {
// 		count++
// 		fmt.Println(count)
// 		// only update the left pointer if
// 		// it's behind the last position of the visited character
// 		l = max(l, index)
// 		fmt.Printf("l:%v\n", l)
// 	}
// 	res = max(res, r-l+1)
// 	hashLen[str[r]] = r + 1
// 	// fmt.Printf("key:%  value:%v\n", str[r], hashLen[str[r]])
// }
// return res
// }

// func max(n, m int) int {
// 	if n > m {
// 		return n
// 	}
// 	return m
// }
