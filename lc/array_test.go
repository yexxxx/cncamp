package lc

import (
	"fmt"
	"github.com/yexxxx/cncamp/lc/ds"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}

	fmt.Println(removeElement(nums, 2))
	fmt.Println(nums)
}

// 26
func removeDuplicates(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	var slow, fast int
	for ; fast < length; fast++ {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}

// 83 删除排序链表中的重复元素
func deleteDuplicates1(head *ds.ListNode) *ds.ListNode {
	slow, fast, dummy := head, head, head

	if head == nil {
		return nil
	}
	for fast != nil {
		if slow.Val != fast.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}
	slow.Next = nil
	return dummy
}

// 83 删除排序链表中的重复元素
func deleteDuplicates2(head *ds.ListNode) *ds.ListNode {
	cur := head

	if head == nil {
		return nil
	}
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

// 27
func removeElement(nums []int, val int) int {
	left := 0
	for _, v := range nums { // v 即 nums[right]
		if v != val {
			nums[left] = v
			left++
		}
	}
	return left
}

// 283. 移动零
func moveZeroes(nums []int) {
	left, length := 0, len(nums)
	for right := 0; right < length; right++ {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}
}

// 167. 两数之和 II - 输入有序数组
func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum < target {
			left++
		} else if sum > target {
			right--
		} else {
			return []int{left + 1, right + 1}
		}
	}
	return nil
}

// 344. 反转字符串
func reverseString(s []byte) {
	for left, right := 0, len(s)-1; left < right; left++ {
		s[left], s[right] = s[right], s[left]
		right--
	}
}

func TestSth(t *testing.T) {
	s := "golang语言"
	for i, s := range s {
		fmt.Printf("%d : %v, %s \n", i, s, string(s))
	}

	arr := []rune(s)
	for i, s := range arr {
		fmt.Printf("%d : %v \n", i, s)
	}

	fmt.Println(len(s))
}

//5. 最长回文子串
//func longestPalindrome(s string) string {
//
//}

//func isPalindrome(s []rune) (count int) {
//
//}

func getListNodeByArr(arr []int) *ds.ListNode {
	dummy := &ds.ListNode{}
	cur := dummy
	for _, v := range arr {
		node := &ds.ListNode{Val: v}
		cur.Next = node
		cur = cur.Next
	}
	return dummy.Next
}
