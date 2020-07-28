package array

func twoSum(nums []int, target int) []int {
	// 设置map的长度以后，内存占用有进一步的缩小，不用重复分配内存了
	m := make(map[int]int, len(nums)-1)

	for i1, v := range nums {
		if i2, ok := m[target-v]; ok {
			return []int{i2, i1}
		}
		m[v] = i1
	}
	return []int{}
}
