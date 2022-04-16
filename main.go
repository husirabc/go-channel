package main

import (
	"bytes"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/toolkits/concurrent/semaphore"
)
//#
//#go-channel-test
//循环读取字符串，在第三秒的时候退出

func main01() {
	atime := time.NewTicker(3 * time.Second)
	ch := make(chan int)
	overch := make(chan int)
	abcstrings := "abcdef"
	go func() {
		defer fmt.Println("follower thread over")
		lenstring := len(abcstrings)
		for i := 0; i < lenstring; i++ {
			select {
			case a := <-ch:
				overch <- a
				return

			default:
				fmt.Println("data:%v", abcstrings[i])
			}
			time.Sleep(1 * time.Second)
		}
	}()

	select {
	case <-atime.C:
		ch <- 1
	}
	<-overch
	time.Sleep(1 * time.Second)
	fmt.Println("main thread over")
	close(ch)
	close(overch)
}

func main02() {
	atime := time.NewTicker(3 * time.Second)
	ch := make(chan int)
	overch := make(chan int)
	abcstrings := "abcdef"

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("follower thread over")
		lenstring := len(abcstrings)
		for i := 0; i < lenstring; i++ {
			select {
			case a := <-ch:
				overch <- a
				return

			default:
				fmt.Println(fmt.Sprintf("data:%v", abcstrings[i]))
			}
			time.Sleep(1 * time.Second)
		}

	}()

	select {
	case <-atime.C:
		ch <- 1
	}
	<-overch
	wg.Wait()
	fmt.Println("main thread over")
	close(ch)
	close(overch)
}

//golang线程控制
func main03() {
	var threadcontrol *semaphore.Semaphore
	threadcontrol = semaphore.NewSemaphore(2)

	go func() {
		threadcontrol.Acquire()

		time.Sleep(5 * time.Second)
		fmt.Println("线程1运行完毕")
		threadcontrol.Release()
	}()

	go func() {
		threadcontrol.Acquire()

		time.Sleep(10 * time.Second)
		fmt.Println("线程2运行完毕")
		threadcontrol.Release()
	}()

	time.Sleep(time.Second)
	go func() {
		threadcontrol.Acquire()

		time.Sleep(1 * time.Second)
		fmt.Println("线程3运行完毕")
		threadcontrol.Release()
	}()

	time.Sleep(20 * time.Second)
	fmt.Println("主线运行完毕")
}

//循环读取字符串，在第三秒的时候退出
//知识点  考察channel的使用

func main011() {

	timeout := time.NewTicker(3 * time.Second)
	S := "abcdefghijklmn"
	seconds3 := make(chan int)
	closechan := make(chan int)
	go func(ss string) {
		lens := len(ss)
		for i := 0; i < lens-1; i++ {
			select {
			case <-seconds3:
				closechan <- 0
				fmt.Println("go func(ss string) over")
				return
			default:
				fmt.Println(fmt.Sprintf("the char:[%d] %c", i, ss[i]))
				time.Sleep(time.Second)
			}

		}
	}(S)

	select {
	case <-timeout.C:
		seconds3 <- 0

	}

	<-closechan
	time.Sleep(time.Second)
	fmt.Println("go main over")
}

func main012() {
	timeout := time.NewTicker(3 * time.Second)
	S := "abcdefghijklmn"
	seconds3 := make(chan int)

	ccc := sync.WaitGroup{}
	go func(ss string) {
		defer ccc.Done()
		ccc.Add(1)
		lens := len(ss)
		for i := 0; i < lens-1; i++ {
			select {
			case <-seconds3:

				fmt.Println("go func(ss string) over")
				return
			default:
				fmt.Println(fmt.Sprintf("the char:[%d] %c", i, ss[i]))
				time.Sleep(time.Second)
			}

		}
	}(S)

	select {
	case <-timeout.C:
		seconds3 <- 0
	}

	ccc.Wait()

	//	<-closechan
	//	time.Sleep(time.Second)
	fmt.Println("go main over")
}

func main013() {
	fmt.Println(getMaxProduct([]int{-2, 3, -2, -3, 4, -3, 2}))
}

func getMaxProduct(arrays []int) int {
	var maxproduct int
	if len(arrays) == 0 {
		maxproduct = 0
	}

	max := make([]int, len(arrays))
	min := make([]int, len(arrays))

	max[0] = arrays[0]
	min[0] = arrays[0]
	maxproduct = max[0]
	for i := 1; i < len(arrays); i++ {
		if arrays[i] > 0 {
			max[i] = max[i-1] * arrays[i]
			min[i] = min[i-1] * arrays[i]
		} else {
			min[i] = max[i-1] * arrays[i]
			max[i] = min[i-1] * arrays[i]
		}

		if max[i] < arrays[i] {
			max[i] = arrays[i]
		}
		if min[i] > arrays[i] {
			min[i] = arrays[i]
		}

		if max[i] > maxproduct {
			maxproduct = max[i]
		}

		fmt.Println(max)
		fmt.Println(min)
	}

	return maxproduct

}

func main014() {
	fmt.Println(convert("abcdefghijklm", 4))

	fmt.Println("2222222:", 1<<3-1)
	fmt.Println("2222222:", math.Pow(3, 2)-2)

}

func getZstring(s string, n int) string {
	sbyte := []byte(s)
	t := 2*n - 2
	news := []byte{}
	for i := 0; i < n; i++ {
		for j := 0; j+i < len(sbyte); j = j + t {
			news = append(news, sbyte[i+j])
			if i > 0 && (j+t-i) < len(s) && i < n-1 {
				news = append(news, sbyte[j+t-i])
			}
			fmt.Println(string(news))
		}
	}
	return string(news)
}

func convert(s string, numRows int) string {
	if s == "" {
		return ""
	}

	if numRows <= 1 || numRows > len(s) {
		return s
	}
	news := []byte{}
	fmt.Println(news)
	t := 2*numRows - 2
	for i := 0; i < numRows; i++ {
		for j := 0; j+i < len(s); j = j + t {
			news = append(news, s[i+j])

			if i > 0 && i < numRows-1 && j-i+t < len(s) {
				news = append(news, s[j-i+t])
			}
			fmt.Println(news)
		}
	}

	return string(news)

}

func myAtoi(s string) int {
	var lowzero bool
	s2 := "-"
	sby := []byte(s2)
	for i := 0; i < len(s); i++ {
		fmt.Println("aaa", s[i])
		if s[i] != 32 {
			if s[i] == sby[0] && (i+1) < len(s) {
				s = s[i+1:]
				fmt.Println("222", s)
				lowzero = true
			} else {
				s = s[i:]
			}
			break
		}
	}

	fmt.Println(s)
	a, err := strconv.Atoi(s)

	if err != nil {
		return 0
	}
	if lowzero {
		return 0 - a
	}
	return a
}

func main015() {
	//fmt.Println(myAtoi("   -42"))
	s := "abc"

	if s[0] == 'a' {
		fmt.Println("ok", s[0], 'a')
	}
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for k := 0; k < len(strs[0]); k++ {
		//for k, sva := range strs[0] {
		for i := 1; i < len(strs)-1; i++ {
			if (len(strs[i]) - 1) < k {
				return strs[0][0:k]
			}
			//	sa :=
			if strs[i][k] != strs[0][k] {
				return strs[0][0:k]
			}
		}
	}
	return strs[0]
}

//第20题
/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
*/

func isValid(s string) bool {
	// if len(s)%2!=0{
	//     return false
	// }

	// cbytes:=[]byte{}
	// cmap:=map[byte]byte{
	//     ')':'(',
	//     ']':'[',
	//     '}':'{',
	// }

	// for i:=0;i<len(s);i++{
	//     if cmap[s[i]]==0{
	//         cbytes=append(cbytes,s[i])
	//     }else{
	//         if len(cbytes)<1{
	//             return false
	//         }

	//         if cmap[s[i]]!=cbytes[len(cbytes)-1]{
	//             return false
	//         }

	//         cbytes=cbytes[:len(cbytes)-1]
	//     }
	// }
	// if len(cbytes)!=0{
	//     return false
	// }
	// return true

	slen := len(s)
	for {

		s = strings.Replace(s, "()", "", -1)
		s = strings.Replace(s, "{}", "", -1)
		s = strings.Replace(s, "[]", "", -1)

		if slen == len(s) {
			return false
		} else {
			slen = len(s)
			if slen == 0 {
				return true
			}
		}
	}
}

//第26题
/*删除有序数组中的重复项
给你一个 升序排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。

由于在某些语言中不能改变数组的长度，所以必须将结果放在数组nums的第一部分。更规范地说，如果在删除重复项之后有 k 个元素，那么 nums 的前 k 个元素应该保存最终结果。

将最终结果插入 nums 的前 k 个位置后返回 k 。

不要使用额外的空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

*/
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var returnlen int = 0
	lastnum := nums[0]
	ok := lastnum - 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == lastnum {
			nums[i] = ok
		} else {
			returnlen++
			lastnum = nums[i]
		}
	}

	for i := 2; i < len(nums); i++ {
		if nums[i] != ok {
			for j := i; j > 1; j-- {
				if nums[j-1] == ok {
					nums[j-1] = nums[j]
					nums[j] = ok
				}
			}
		}
	}

	return returnlen + 1
}

func main27() {
	arraya := []int{3, 2, 2, 3}
	fmt.Println(removeElement(arraya, 3))
}

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	lval := 0
	lnums := 0
	for i := 0; lnums < len(nums); {
		lnums++
		fmt.Println(nums[i])
		if nums[i] == val {
			for j := i; j < len(nums)-1; j++ {
				nums[j] = nums[j+1]
			}
			continue
		}
		lval++
		i++
	}
	fmt.Println(nums)
	nums = nums[:lval]
	fmt.Println(nums)

	return lval
}

/*
21 合并两个有序链表
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	var va1, va2 int
	var middleList *ListNode = &ListNode{} //********
	returnList := middleList
	for list1 != nil && list2 != nil {
		va1 = list1.Val
		va2 = list2.Val

		if va1 <= va2 {
			middleList.Val = list1.Val

			list1 = list1.Next
		} else {
			middleList.Val = list2.Val
			list2 = list2.Next
		}
		middleList.Next = &ListNode{}
		middleList = middleList.Next
	}
	if list1 != nil {
		middleList.Val = list1.Val
		middleList.Next = list1.Next
	}
	if list2 != nil {
		middleList.Val = list2.Val
		middleList.Next = list2.Next
	}
	return returnList
}

/*
28 实现 strStr()
实现 strStr() 函数。

给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回  -1 。

说明：

当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与 C 语言的 strstr() 以及 Java 的 indexOf() 定义相符。

*/

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] && (i+len(needle)) <= len(haystack) {
			hhay := haystack[i : i+len(needle)]
			if string(hhay) == needle {
				return i
			}
		}
	}
	return -1
}

/*
35 搜索插入位置e

给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

请必须使用时间复杂度为 O(log n) 的算法。
*/

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if nums[0] > target {
		return 0
	}
	if nums[len(nums)-1] < target {
		return len(nums)
	}
	var numint [2]int = [2]int{0, len(nums) - 1}
	for {
		// numint++
		n := (len(nums) - 1) / 2
		if len(nums) == 2 {
			n = 1
		}
		if nums[n] < target {
			if n == 0 {

				fmt.Println("return", numint[0])
				return numint[0] + 1
			}
			// numint[0]=n+1

			numint[0] = numint[0] + n
			fmt.Println(numint[0])
			// numint[1]=len(nums)-1
			nums = nums[n:]
			// if n
		} else if nums[n] > target {
			if n == 0 {
				return numint[1] + 1
			}
			// numint[0]=0
			numint[1] = numint[1] - n
			nums = nums[:n]
		} else {
			x := numint[0] + n
			return x
		}
	}
}

func ttt() [][]int {
	aaa := [][]int{{0, 0, 0}}
	aaa = [][]int{}
	return aaa
}

func threeSum(nums []int) [][]int {
	lennums := len(nums)
	if lennums < 3 {
		return nil
	}
	sort.Ints(nums)
	fmt.Println(nums)
	//fmt.Println(nums)
	threearrays := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		target := 0 - nums[i]
		fmt.Println("target:", target)
		if i != 0 && nums[i-1] == nums[i] {
			continue
		}

		twoarrays := twonum(nums[i:], target)
		threearrays = append(threearrays, twoarrays...)
	}
	return threearrays
}

func twonum(numsafter []int, target int) [][]int {
	fmt.Println(numsafter)
	twoarrays := make([][]int, 0)
	i, j := 1, len(numsafter)-1
	for i < j {
		m := numsafter[i] + numsafter[j]
		//  twoarrays=append(twoarrays,[]int{numsafter[i],numsafter[j]})
		if m == target {
			twoarrays = append(twoarrays, []int{numsafter[0], numsafter[i], numsafter[j]})
			// 去重
			for i < j {
				i++
				if numsafter[i-1] != numsafter[i] {
					break
				}
			}
			for i < j {
				j--
				if numsafter[j+1] != numsafter[j] {
					break
				}
			}
		}
		if m < target {
			i++
		}
		if m > target {
			j--
		}
	}
	return twoarrays
}

func threeSum2(nums []int) [][]int {

	if len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	//fmt.Println(nums)
	//fmt.Println(nums)
	threearrays := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		target := -nums[i]
		//	fmt.Println("target:", target)
		if i != 0 && nums[i-1] == nums[i] {
			continue
		}

		//twoarrays := twonum(nums[i:], target)

		ii, j := i+1, len(nums)-1
		for ii < j {
			var m = nums[ii] + nums[j]

			if m == target {
				threearrays = append(threearrays, []int{nums[i], nums[ii], nums[j]})
				// 去重
				for ii < j {
					ii++
					if nums[ii-1] != nums[ii] {
						break
					}
				}
				for ii < j {
					j--
					if nums[j+1] != nums[j] {
						break
					}
				}
			}
			if m < target {
				ii++
			}
			if m > target {
				j--
			}
		}
	}
	return threearrays
}

func main222() {
	fmt.Println(fourSum([]int{-2, -1, -1, 1, 1, 2, 2}, 0))
}

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}

	sort.Ints(nums)

	var newnums = make([][]int, 0)

	for i := 0; i < len(nums)-3; i++ {
		fmt.Println("iiii:", nums[i])
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		onetarget := target - nums[i]
		for j := i + 1; j < len(nums)-2; j++ {
			// if j != 1 && nums[j] == nums[j-1] {
			// 	continue
			// }
			twotarget := onetarget - nums[j]
			if nums[i] == -1 {
				fmt.Println("jjj:", nums[j], twotarget)
			}

			left, right := j+1, len(nums)-1
			for left < right {
				m := nums[left] + nums[right]
				if m == twotarget {
					newnums = append(newnums, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right {
						left++
						if nums[left] != nums[left-1] {
							break
						}
					}
					for left < right {
						right--
						if nums[right] != nums[right+1] {
							break
						}
					}

				}
				if m > twotarget {
					right--
				}
				if m < twotarget {
					left++
				}
			}
		}
	}
	return newnums
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	l := l3
	var nextvalue = 0
	//  var sum=0
	var nextplus = 0
	for l1 != nil && l2 != nil {
		value1 := l1.Val
		value2 := l2.Val

		nextvalue = value1 + value2 + nextplus
		if nextvalue >= 10 {
			l3.Val = nextvalue % 10
			nextplus = 1
		} else {
			l3.Val = nextvalue
		}
		l1 = l1.Next
		l2 = l2.Next

		l3.Next = &ListNode{}
		l3 = l3.Next
	}
	if l1 != nil {
		l3 = l1
	} else {
		l3 = l2
	}
	for nextplus != 0 {
		l3.Val = l3.Val + nextplus
		if l3.Val >= 10 {
			l3.Val = l3.Val % 10
			nextplus = 1
		}
		l3 = l3.Next
	}
	return l
}

func addBinary(a string, b string) string {
	ss := ""
	i, j := len(a)-1, len(b)-1
	var plus byte = 0
	for i >= 0 || j >= 0 {

		value1, value2 := byte(0), byte(0)
		if i >= 0 {
			value1 = a[i]
		}
		if j >= 0 {
			value2 = b[j]
		}
		sum := value1 + value2 - '0' - '0' + plus
		if sum >= 2 {
			plus = 1
			sum = sum - 2 + '0'
		} else {
			sum = sum + '0'
		}
		ss = string([]byte{sum}) + ss
		i--
		j--
	}
	if plus == 1 {
		ss = string([]byte{'1'}) + ss
	}
	return ss
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	maxint := []int{0, 0}
	currentints := []int{0, 0}
	maps := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if value, ok := maps[s[i]]; ok {

			for maps[s[value+1]] != value+1 {

				if value+1 > i {
					break
				}
				value++
			}
			currentints[0] = value + 1
			currentints[1] = i

			if maxint[1]-maxint[0] < currentints[1]-currentints[0] {
				maxint[0] = currentints[0]
				maxint[1] = currentints[1]
			}
			maps[s[i]] = i
		} else {
			maps[s[i]] = i
			currentints[1] = i
			if maxint[1]-maxint[0] < currentints[1]-currentints[0] {
				maxint[0] = currentints[0]
				maxint[1] = currentints[1]
			}
		}
	}
	return maxint[1] - maxint[0] + 1
}

func main023() {
	n := 5
	time.Sleep(time.Duration(n) * time.Second)
	now := time.Now()
	fmt.Println(now.Format("02/1/2006 15:04"))
	fmt.Println(now.Format("2006/1/02 15:04"))
	fmt.Println(now.Format("2006/1/02"))

	var a int = 5
	var p = &a
	fmt.Println(*p)

	for i := 0; i < 100; i++ {
	}

	//	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{1, 3, 4}))
}

//二分法查找中位数
func findMedianSortedArrays02(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)

	if len(nums1) > len(nums2) {
		return findMedianSortedArrays02(nums2, nums1)
	}

	imin, imax, leftlen := 0, m, (m+n+1)/2
	for imin <= imax {
		i := imin + (imax-imin)/2
		j := leftlen - i

		if i > imin && nums1[i-1] > nums2[j] {
			imax = i - 1
		} else if nums2[j-1] > nums1[i] {
			imin = i + 1
		}
	}

	return 0.0
}

// func main14() {
// 	n := 4
// 	fmt.Println(PrintOpenClose(n))

// 	//	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{1, 3, 4}))
// }

func mainaa() {
	// aa := make(map[int]int)
	// test(aa)
	// fmt.Println(&aa)
	// fmt.Println(fmt.Sprintf("%p", aa))

	// bb := []int{1, 3, 4}
	// test2(bb)
	// fmt.Println(&bb)
	// fmt.Println(fmt.Sprintf("%p", bb))

	//	spew.Dump(aa)

	//n := 3
	//fmt.Println(generateParenthesis(n))
	bytes.NewBuffer([]byte("aa"))

	//fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{1, 3, 4}))

}

func removeDuplicates2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var num = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			if i > num {
				nums[num] = nums[i]
			}
			num++
		}
	}
	return num
}

func removeElement2(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	left, right := 0, len(nums)-1
	numok := len(nums)
	for left < right {
		if nums[left] == val {
			numok--
			//left++
			for left < right {
				if nums[right] != val {
					nums[left] = nums[right]
					right--
					break
				}
				numok--
				right--
			}

		}
		left++
	}
	return numok
}

func searchInsert2(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	left, right := 0, len(nums)-1
	if target > nums[right] {
		return right + 1
	}
	if target < nums[left] {
		return 0
	}
	for left < right {

		i := (right-left+1)/2 + left

		if nums[i] == target {
			return i
		} else if nums[i] > target {
			right = i
		} else {
			left = i
		}
		if left == right {
			if nums[left] == target {
				return left
			}
			return left + 1
		}
	}
	return 0
}
