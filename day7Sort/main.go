package main

import "fmt"

func main() {
	sort := InsertSort([]int{5, 2, 4, 6, 1, 3})
	fmt.Println(sort)

	bobSort := BobSort([]int{5, 2, 4, 6, 1, 3})
	fmt.Println(bobSort)

	selectSort := SelectSort([]int{5, 2, 4, 6, 1, 3})
	fmt.Println(selectSort)

	mergeSort := MergeSort([]int{5, 2, 4, 6, 1, 3})
	fmt.Println(mergeSort)

	quickSort := QuickSort([]int{6, 4, 8, 2, 3, 1, 5, 7, 9})
	fmt.Println(quickSort)
}

// QuickSort 快排 先拿到再排序 先把拿到的值放在某个位置
//这个位置的左边通通小于该值 这个位置的右边通通大于该值
func QuickSort(ints []int) []int {

	var travel func(ints []int)
	travel = func(ints []int) {
		if len(ints) == 0 || len(ints) == 1 {
			return
		}
		target := ints[0]
		left, right := 0, len(ints)-1
		for left != right {
			for ints[right] >= target && left < right {
				right--
			}
			if left != right {
				ints[left] = ints[right]
				left++
			}
			for ints[left] <= target && left < right {
				left++
			}
			if left != right {
				ints[right] = ints[left]
				right--
			}
		}
		ints[right] = target

		travel(ints[:left])
		travel(ints[left+1:])
	}

	travel(ints)
	return ints
}

// MergeSort 归并排序 先拆 拆成一个之后再排 排的时候有序排列
//保证最后排出来的就是有序的
func MergeSort(ints []int) []int {
	var merge func(a []int, b []int) []int
	merge = func(a []int, b []int) []int {
		if len(a) == 0 && len(b) == 0 {
			return []int{}
		}

		ans := make([]int, 0, len(a)+len(b))
		la := len(a)
		lb := len(b)
		indexa := 0
		indexb := 0
		for indexa < la || indexb < lb {
			if indexa == la {
				ans = append(ans, b[indexb])
				indexb++
				continue
			}

			if indexb == lb {
				ans = append(ans, a[indexa])
				indexa++
				continue
			}

			if a[indexa] < b[indexb] {
				ans = append(ans, a[indexa])
				indexa++
			} else {
				ans = append(ans, b[indexb])
				indexb++
			}
		}
		return ans
	}

	var travel func([]int) []int
	travel = func(ints []int) []int {
		mid := len(ints) / 2
		if mid == 0 {
			return ints
		}

		left := travel(ints[:mid])
		right := travel(ints[mid:])
		i := merge(left, right)
		return i
	}

	return travel(ints)
}

// SelectSort 选择排序  减少交换次数
//每次寻找最小的元素 然后放置到相应的位置
func SelectSort(ints []int) []int {
	for i := 0; i < len(ints); i++ {
		min := i
		for j := i + 1; j < len(ints); j++ {
			if ints[min] > ints[j] {
				min = j
			}
		}
		ints[i], ints[min] = ints[min], ints[i]
	}
	return ints
}

// BobSort 冒泡排序 相邻两个元素两两交换
//每轮就能得到当前轮数的最大值
func BobSort(ints []int) []int {
	count := 0
	for count != len(ints) {
		for i := 0; i < len(ints)-1-count; i++ {
			if ints[i] > ints[i+1] {
				ints[i], ints[i+1] = ints[i+1], ints[i]
			}
		}
		count++
	}
	return ints
}

// InsertSort 插入排序
// 从后往前 依次判断 如果发现当前的target大于等于前面的 就break
// 从后往前 依次判断 如果发现当前的target小于前面的 就说明要插到前面去 就把前面的集体后移
func InsertSort(ints []int) []int {
	for i := 1; i < len(ints); i++ {
		target := ints[i]
		for j := i - 1; j >= 0; j-- {
			if ints[j] > target {
				ints[j], ints[j+1] = target, ints[j]
			} else {
				break
			}
		}
	}
	return ints
}
