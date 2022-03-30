package conversion

import "strconv"

func Int64toA(n int64) string {
	return strconv.FormatInt(n, 10)
}

func AtoInt64Must(a string) int64 {
	i, _ := strconv.ParseInt(a, 10, 64)
	return i
}
