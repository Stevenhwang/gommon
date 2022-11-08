package slice

// FindValInStringSlice 查询val是否在 string slice 中
func FindValInStringSlice(slice []string, val string) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

// RemoveDupInStringSlice 去掉 string slice 中的重复元素
func RemoveDupInStringSlice(slice []string) []string {
	result := []string{}
	tempMap := map[string]byte{} // 存放不重复主键
	for _, e := range slice {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}
