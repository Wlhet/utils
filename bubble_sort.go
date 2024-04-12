package utils

// 泛型冒泡排序
func BubbleSort[T any](arr []T, less func(T, T) bool) {
	/*
			intArr := []int{64, 34, 25, 12, 22, 11, 90}
		fmt.Printf("Before sorting: %v\n", intArr)

		// 使用泛型冒泡排序对整型切片进行排序
		BubbleSort(intArr, func(a, b int) bool {
			return a < b
		})

		fmt.Printf("After sorting: %v\n", intArr)

		strArr := []string{"apple", "orange", "banana", "pear", "grapes"}
		fmt.Printf("Before sorting: %v\n", strArr)

		// 使用泛型冒泡排序对字符串切片进行排序
		BubbleSort(strArr, func(a, b string) bool {
			return len(a) < len(b)
		})

		fmt.Printf("After sorting: %v\n", strArr)
	*/
	arrLen := len(arr)
	for i := 0; i < arrLen-1; i++ {
		for j := 0; j < arrLen-i-1; j++ {
			if less(arr[j+1], arr[j]) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
