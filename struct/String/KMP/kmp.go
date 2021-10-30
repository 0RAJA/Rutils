package KMP

// Kmp s1为主串,s2为模板串,如果s2在s1中出现,返回s1中第一个匹配的字符下标,否则返回-1
func Kmp(s1, s2 string) int {
	if len(s1) < len(s2) {
		return -1
	}
	next := GetNext(s2)
	for i, j := 0, 0; i < len(s1); {
		if j == -1 || s1[i] == s2[j] {
			i++
			j++
		} else {
			j = next[j]
		}
		if j == len(s2) {
			return i - j
		}
	}
	return -1
}

// GetNext 返回模板串s的next数组
func GetNext(s string) []int {
	next := make([]int, len(s))
	next[0] = -1
	for i, j := 0, -1; i < len(s)-1; {
		if j == -1 || s[i] == s[j] {
			i++
			j++
			if s[i] != s[j] {
				next[i] = j
			} else {
				next[i] = next[j]
			}
		} else {
			j = next[j]
		}
	}
	return next
}


