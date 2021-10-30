package Subsequence

// Subsequence 返回s1和s2的最长子序列长度
func Subsequence(s1, s2 string) int {
	dp := make([][]int, len(s1)+1) //行是len(s1)+1
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(s2)+1) //列是len(s2)+1
	}
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = func() int {
					if dp[i-1][j] > dp[i][j-1] {
						return dp[i-1][j]
					}
					return dp[i][j-1]
				}()
			}
		}
	}
	return dp[len(s1)][len(s2)]
}
