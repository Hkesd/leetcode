package main

func lengthOfLongestSubstring(s string) int {
	stringMap := make(map[rune]int, len(s))

	ans, start := 0, 0

	for k, v := range s {
		if lastIndex, ok := stringMap[v]; ok {
			start = maxInt(start, lastIndex)
		}
		ans = maxInt(ans, k - start + 1)

		stringMap[v] = k + 1
	}

	return ans
}

func maxInt(x, y int) int {
	if x < y {
		return y
	}

	return  x
}

//func main()  {
//	s := "abcabcbb"
//
//	re := lengthOfLongestSubstring(s)
//
//	fmt.Println(re)
//}
