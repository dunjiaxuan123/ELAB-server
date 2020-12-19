package api

import "strconv"

func QuickSort(arr map[int]map[string]string, first int, last int) map[int]map[string]string {
	flag := first
	left := first
	right := last

	if first >= last {
		return arr
	}
	// 将大于arr[flag]的都放在右边，小于的，都放在左边
	for first < last {
		// 如果flag从左边开始，那么是必须先从有右边开始比较，也就是先在右边找比flag小的
		intFlag, _ := strconv.ParseInt(arr[flag]["timestamp"], 10, 64)
		intFirst, _ := strconv.ParseInt(arr[first]["timestamp"], 10, 64)
		intLast, _ := strconv.ParseInt(arr[last]["timestamp"], 10, 64)
		for first < last {
			if intLast >= intFlag {
				last--
				continue
			}
			// 交换数据
			arr[last], arr[flag] = arr[flag], arr[last]
			flag = last
			break
		}
		for first < last {
			if intFirst <= intFlag {
				first++
				continue
			}
			arr[first], arr[flag] = arr[flag], arr[first]
			flag = first
			break
		}
	}

	QuickSort(arr, left, flag-1)
	QuickSort(arr, flag+1, right)
	return arr
}
