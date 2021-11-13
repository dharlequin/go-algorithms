package algorithms

func NumInList(list []int, num int) bool {
	if len(list) == 0 {
		return false
	}

	for _, number := range list {
		if number == num {
			return true
		}
	}

	return false
}
