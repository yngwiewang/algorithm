package common

// MaxInt 取得两个整数的较大值
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// DistanceInt 计算两个整数值的距离（差的绝对值）
func DistanceInt(a, b int) int {
	if a-b < 0 {
		return -1 * (a - b)
	}
	return a - b
}
