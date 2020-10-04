package common

// MaxInt 取得两个整数的较大值
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// MinInt 取得两个整数的较小值
func MinInt(a, b int) int {
	if a < b {
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

// EqualInts 比较两个整数切片是否相等
func EqualInts(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// MinIntInSlice 获得整形切片的最小值
func MinIntInSlice(n []int) int {
	min := n[0]
	for _, v := range n {
		if v < min {
			min = v
		}
	}
	return min
}
