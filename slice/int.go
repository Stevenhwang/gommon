package slice

// FindValInIntSlice 查询 val 是否在 int slice 中
func FindValInIntSlice(slice []int, val int) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

// RemoveDupInIntSlice 去掉 int slice 中的重复元素
func RemoveDupInIntSlice(slice []int) []int {
	result := []int{}
	tempMap := map[int]byte{} // 存放不重复主键
	for _, e := range slice {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}
