package helper

func RemoveDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func RemoveDuplicateInt(intSlice []int) []int {
	allKeys := make(map[int]bool)
	list := []int{}
	for _, item := range intSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func IncludesStr(strSlice []string, finds []string) bool {
	check := false

	for _, value := range strSlice {
		for _, find := range finds {
			if value == find {
				check = true
				break
			}
		}
	}

	return check
}

func IncludesInt(intSlice []int, finds []int) bool {
	check := false

	for _, value := range intSlice {
		for _, find := range finds {
			if value == find {
				check = true
				break
			}
		}
	}

	return check
}

func SliceIndexStr(strSlice []string, find string) int {
	index := -1

	for key, value := range strSlice {
		if value == find {
			index = key
			break
		}
	}

	return index
}

func SliceIndexInt(strSlice []int, find int) int {
	index := -1

	for key, value := range strSlice {
		if value == find {
			index = key
			break
		}
	}

	return index
}
