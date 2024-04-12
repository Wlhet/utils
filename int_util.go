package utils

func FindDifferenceWithInt(A, B []int) (ANotInB []int, BNotInA []int, Common []int) {
	return SliceDifference(A, B)
}

// Min golang min int
func Min(first int, args ...int) int {
	for _, v := range args {
		if first > v {
			first = v
		}
	}
	return first
}
