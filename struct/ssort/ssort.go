package ssort

// BubbleSort 冒泡排序 n^2
func BubbleSort(list []int, cmp func(i, j int) bool) {
	flag := true
	for i := 0; i < len(list); i++ {
		flag = true
		for j := len(list) - 2; j >= i; j-- { // 从后往前冒泡
			if !cmp(j, j+1) {
				flag = false
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
		if flag == true {
			return
		}
	}
}

// SelectSort 选择排序 n^2
func SelectSort(list []int, cmp func(i, j int) bool) {
	n := len(list)
	for i := 0; i < n/2; i++ {
		minIndex, maxIndex := i, i
		for j := i + 1; j < n-i; j++ {
			if cmp(maxIndex, j) {
				maxIndex = j
				continue
			}
			if !cmp(minIndex, j) {
				minIndex = j
			}
		}
		if maxIndex == i && minIndex != n-i-1 {
			// 如果最大值是开头的元素，而最小值不是最尾的元素
			// 先将最大值和最尾的元素交换
			list[n-i-1], list[maxIndex] = list[maxIndex], list[n-i-1]
			// 然后最小的元素放在最开头
			list[i], list[minIndex] = list[minIndex], list[i]
		} else if maxIndex == i && minIndex == n-i-1 {
			// 如果最大值在开头，最小值在结尾，直接交换
			list[minIndex], list[maxIndex] = list[maxIndex], list[minIndex]
		} else {
			// 否则先将最小值放在开头，再将最大值放在结尾
			list[i], list[minIndex] = list[minIndex], list[i]
			list[n-i-1], list[maxIndex] = list[maxIndex], list[n-i-1]
		}
	}
}

// InsertSort 插入排序 n^2
func InsertSort(list []int, cmp func(num1, num2 int) bool) {
	n := len(list)
	// 进行 N-1 轮迭代
	for i := 1; i <= n-1; i++ {
		deal := list[i] // 待排序的数
		j := i - 1      // 待排序的数左边的第一个数的位置
		// 如果第一次比较，比左边的已排好序的第一个数小，那么进入处理
		if cmp(deal, list[j]) {
			// 一直往左边找，比待排序大的数都往后挪，腾空位给待排序插入
			for ; j >= 0 && cmp(deal, list[j]); j-- {
				list[j+1] = list[j] // 某数后移，给待排序留空位
			}
			list[j+1] = deal // 结束了，待排序的数插入空位
		}
	}
}

// ShellSort 希尔排序
func ShellSort(list []int, cmp func(num1, num2 int) bool) {
	// 数组长度
	n := len(list)
	// 每次减半，直到步长为 1
	for step := n / 2; step >= 1; step /= 2 {
		// 开始插入排序，每一轮的步长为 step
		for i := step; i < n; i += step {
			for j := i - step; j >= 0; j -= step {
				// 满足插入那么交换元素
				if cmp(list[j+step], list[j]) {
					list[j], list[j+step] = list[j+step], list[j]
					continue
				}
				break
			}
		}
	}
}

// MergeSort 归并排序 nlogn
func MergeSort(nums []int, l, r int) {
	// 至少有两个元素,进入排序
	if r-l > 1 {
		mid := l + (r-l)/2
		MergeSort(nums, l, mid)
		MergeSort(nums, mid, r)
		merge(nums, l, mid, r)
	}
}

func merge(nums []int, begin, mid, end int) {
	tmp := make([]int, 0, end-begin+1)
	l, r := begin, mid
	for l < mid && r < end {
		if nums[r] < nums[l] {
			tmp = append(tmp, nums[r])
			r++
		} else {
			tmp = append(tmp, nums[l])
			l++
		}
	}
	tmp = append(tmp, nums[l:mid]...)
	tmp = append(tmp, nums[r:end]...)
	for i := range tmp {
		nums[begin+i] = tmp[i]
	}
}

// MergeSort2 非递归
func MergeSort2(array []int, begin int, end int) {
	// 步长从1开始
	step := 1
	for end-begin > step {
		for i := begin; i < end; i += step << 1 { // 每次找到两个合并部分
			lo := i
			mid := lo + step
			ro := mid + step
			if mid > end { // 不到一个的范围就break
				break
			}
			if ro > end {
				ro = end
			}
			merge(array, lo, mid, ro)
		}
		step <<= 1
	}
}

// QSort 快排 nlogn cmp(a,b) 比较的是值
func QSort(array []int, cmp func(num1, num2 int) bool) {
	low, high := 0, len(array)-1
	i, j := low, high
	if low >= high {
		return
	}
	k := array[low]
	for low < high {
		for ; low < high && !cmp(array[high], k); high-- {
		}
		if cmp(array[high], k) {
			array[low] = array[high]
		}
		for ; low < high && !cmp(k, array[low]); low++ {
		}
		if cmp(k, array[low]) {
			array[high] = array[low]
		}
	}
	array[low] = k
	QSort(array[i:low+1], cmp)
	QSort(array[low+1:j+1], cmp)
}

// QSortStable 稳定快排 cmp == true 从小到大
func QSortStable(array []int, b bool) {
	if len(array) <= 1 {
		return
	}
	length := len(array)
	arr1 := make([]int, 0, length)
	arr2 := make([]int, 0, length)
	k := array[0]
	for idx := 0; idx < length; idx++ {
		if b && array[idx] >= k {
			arr2 = append(arr2, array[idx])
		} else {
			arr1 = append(arr1, array[idx])
		}
	}
	mid := len(arr1)
	arr1 = append(arr1, arr2...)
	for t := 0; t < length; t++ {
		array[t] = arr1[t]
	}
	if mid > 0 {
		QSortStable(array[:mid], b)
	}
	if mid < len(array)-1 {
		QSortStable(array[mid+1:], b)
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// SortList 给出链表的第一个有效节点和最后一个有效节点的下一个节点
func SortList(head, tail *ListNode) *ListNode {
	// 0
	if head == nil {
		return head
	}
	// 1
	if head.Next == tail {
		head.Next = nil
		return head
	}
	// >=2 快慢指针找中点
	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	mid := slow
	// merge
	return mergeList(SortList(head, mid), SortList(mid, tail))
}

func mergeList(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	// 合并两个有序链表
	t, t1, t2 := dummyHead, head1, head2
	for t1 != nil && t2 != nil {
		if t1.Val <= t2.Val {
			t.Next = t1
			t1 = t1.Next
		} else {
			t.Next = t2
			t2 = t2.Next
		}
		t = t.Next
	}
	if t1 != nil {
		t.Next = t1
	} else {
		t.Next = t2
	}
	return dummyHead.Next
}

// SearchMiddle 寻找中位数
func SearchMiddle(arr []int, start, end int) int {
	k := arr[start]
	i, j := start, end
	if start >= end {
		return arr[start]
	}
	for i < j {
		for i < end && arr[i] <= k {
			i++
		}
		for j > start && arr[j] >= k {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		} else {
			break
		}
	}
	if k > arr[j] {
		arr[start] = arr[j]
		arr[j] = k
	}
	if j == (len(arr)-1)/2 {
		return arr[j]
	} else if j < (len(arr)-1)/2 {
		return SearchMiddle(arr, j+1, end)
	} else {
		return SearchMiddle(arr, start, j-1)
	}
}
