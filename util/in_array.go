package util

func InTxTypeArray(target int32, array []int32) bool {
	for _, item := range array {
		if item == target {
			return true
		}
	}

	return false
}
