package lc

import (
	"fmt"
	"github.com/yexxxx/cncamp/lc/ds"
	"testing"
)

func TestDetectCycle(t *testing.T) {
	head := getCycleNode()
	//pos := detectCycle(head)
	pos := detectCycleWithSet(head)
	fmt.Println(pos.Val)
}

func TestHasCycle(t *testing.T) {
	isCycle, meetPos := hasCycle(getCycleNode())
	fmt.Println(isCycle, meetPos.Val)
}

func TestMergeTwoLists(t *testing.T) {
	lists := []*ds.ListNode{
		ds.GetListNode(),
		ds.GetListNode(),
	}
	result := mergeKLists(lists)

	fmt.Println(result)
}

// 23
func mergeKLists(lists []*ds.ListNode) *ds.ListNode {
	if lists == nil || len(lists) <= 0 {
		return nil
	}
	result := lists[0]
	for i := 1; i < len(lists); i++ {
		result = mergeTwoLists(result, lists[i])
	}
	return result
}

// 23
func mergeKLists2(lists []*ds.ListNode) *ds.ListNode {
	length := len(lists)

	if lists == nil || length < 1 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}
	num := length >> 1
	left := mergeKLists2(lists[:num])
	right := mergeKLists2(lists[num:])
	return mergeTwoLists(left, right)
}

// 21
func mergeTwoLists(list1, list2 *ds.ListNode) *ds.ListNode {
	dummy := &ds.ListNode{}
	pre := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			pre.Next = list1
			list1 = list1.Next
		} else {
			pre.Next = list2
			list2 = list2.Next
		}
		pre = pre.Next
	}
	if list1 != nil {
		pre.Next = list1
	} else {
		pre.Next = list2
	}
	return dummy.Next
}

// 86
func TestPartition(t *testing.T) {
	head, x := getListNodeWithInt()

	fmt.Println(head)
	left := &ds.ListNode{}
	right := &ds.ListNode{}
	leftHead := left
	rightHead := right

	for head != nil {
		if head.Val < x {
			left.Next = head
			left = left.Next
		} else {
			right.Next = head
			right = right.Next
		}
		head = head.Next
		//temp := head.Next
		//head.Next = nil
		//head = temp
	}
	right.Next = nil
	left.Next = rightHead.Next
	fmt.Println(leftHead.Next)
}

func TestListNode_String(t *testing.T) {
	head, _ := getListNodeWithInt()
	fmt.Println(head)
}

func TestIntersectionNode(t *testing.T) {
	headA, headB := getIntersectionNodes()

	node := getIntersectionNodeWithSets(headA, headB)
	if node != nil {
		fmt.Println(node.Val)
	} else {
		fmt.Println("nil")
	}
}

// 160
func getIntersectionNode(headA, headB *ds.ListNode) *ds.ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	first, second := headA, headB
	flagA, flagB := false, false
	for first != second {
		if first == nil && !flagA {
			first = headB
			flagA = true
		}
		if second == nil && !flagB {
			second = headA
			flagB = true
		}

		if first == nil || second == nil {
			return nil
		}
		first = first.Next
		second = second.Next
	}
	return first
}

// 160
func getIntersectionNodeWithSets(headA, headB *ds.ListNode) *ds.ListNode {
	seen := map[*ds.ListNode]struct{}{}
	flag := false

	for headA != nil {
		if _, ok := seen[headA]; ok {
			return headA
		}
		seen[headA] = struct{}{}
		if headA.Next == nil && !flag {
			headA = headB
			flag = true
		} else {
			headA = headA.Next
		}

	}
	return nil
}

func TestRemoveNthFromEnd(t *testing.T) {
	node, _ := getListNodeWithInt()
	fmt.Println(node)
	removeNthFromEnd(node, 3)
	fmt.Println(node)
}

// 19
func removeNthFromEnd(head *ds.ListNode, n int) *ds.ListNode {

	dummy := &ds.ListNode{0, head}
	//倒数n+1
	var nNode = dummy
	for head != nil {
		head = head.Next
		n--
		if n < 0 {
			nNode = nNode.Next
		}
	}
	nNode.Next = nNode.Next.Next
	return dummy.Next
}

//func getListNode() *ds.ListNode {
//	node1 := &ds.ListNode{Val: 1}
//	node2 := &ds.ListNode{Val: 2}
//	node3 := &ds.ListNode{Val: 3}
//	node4 := &ds.ListNode{Val: 4}
//	node5 := &ds.ListNode{Val: 5}
//	node6 := &ds.ListNode{Val: 6}
//
//	node1.Next = node2
//	node2.Next = node3
//	node3.Next = node4
//	node4.Next = node5
//	node5.Next = node6
//
//	return node1
//}

func getListNodeWithInt() (*ds.ListNode, int) {
	node1 := &ds.ListNode{Val: 1}
	node2 := &ds.ListNode{Val: 4}
	node3 := &ds.ListNode{Val: 3}
	node4 := &ds.ListNode{Val: 2}
	node5 := &ds.ListNode{Val: 5}
	node6 := &ds.ListNode{Val: 2}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6

	return node1, 3
}

func getCycleNode() *ds.ListNode {
	node1 := &ds.ListNode{Val: 1}
	node2 := &ds.ListNode{Val: 4}
	node3 := &ds.ListNode{Val: 3}
	node4 := &ds.ListNode{Val: 2}
	node5 := &ds.ListNode{Val: 5}
	node6 := &ds.ListNode{Val: 2}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6

	//node3 is the circle pos
	node6.Next = node3

	return node1
}

func getIntersectionNodes() (headA, headB *ds.ListNode) {
	nodeA1 := ds.ListNode{Val: 4}
	nodeA2 := ds.ListNode{Val: 1}

	nodeB1 := ds.ListNode{Val: 5}
	nodeB2 := ds.ListNode{Val: 6}
	nodeB3 := ds.ListNode{Val: 1}

	node1 := ds.ListNode{Val: 8}
	node2 := ds.ListNode{Val: 4}
	node3 := ds.ListNode{Val: 5}
	node1.Next = &node2
	node2.Next = &node3

	//make node1
	nodeA1.Next = &nodeA2
	nodeA2.Next = &node1

	//make node2
	nodeB1.Next = &nodeB2
	nodeB2.Next = &nodeB3
	nodeB3.Next = &node1

	return &nodeA1, &nodeB1
}

func detectCycleWithSet(head *ds.ListNode) *ds.ListNode {
	seen := map[*ds.ListNode]struct{}{}
	for head != nil {
		if _, ok := seen[head]; ok {
			return head
		}
		seen[head] = struct{}{}
		head = head.Next
	}
	return nil
}

func detectCycle(head *ds.ListNode) *ds.ListNode {
	isCycle, fast := hasCycle(head)
	if isCycle {
		slow := head
		for slow != fast {
			slow = slow.Next
			fast = fast.Next
		}
		return slow
	}
	return nil
}

func hasCycle(head *ds.ListNode) (isCycle bool, meetPos *ds.ListNode) {
	if head == nil || head.Next == nil {
		return false, nil
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true, slow
		}
	}
	return false, nil
}
