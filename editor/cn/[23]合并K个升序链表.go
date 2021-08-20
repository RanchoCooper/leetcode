package main

import "container/heap"

//给你一个链表数组，每个链表都已经按升序排列。
//
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。 
//
// 
//
// 示例 1： 
//
// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
//输出：[1,1,2,3,4,4,5,6]
//解释：链表数组如下：
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//将它们合并到一个有序链表中得到。
//1->1->2->3->4->4->5->6
// 
//
// 示例 2： 
//
// 输入：lists = []
//输出：[]
// 
//
// 示例 3： 
//
// 输入：lists = [[]]
//输出：[]
// 
//
// 
//
// 提示： 
//
// 
// k == lists.length 
// 0 <= k <= 10^4 
// 0 <= lists[i].length <= 500 
// -10^4 <= lists[i][j] <= 10^4 
// lists[i] 按 升序 排列 
// lists[i].length 的总和不超过 10^4 
// 
// Related Topics 链表 分治 堆（优先队列） 归并排序 
// 👍 1365 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type NodeSlice []*ListNode

func (n NodeSlice) Len() int {
	return len(n)
}

func (n NodeSlice) Less(i, j int) bool {
	return n[i].Val < n[j].Val
}

func (n *NodeSlice) Swap(i, j int) {
	(*n)[i], (*n)[j] = (*n)[j], (*n)[i]
}

func (n *NodeSlice) Push(v interface{}) {
	node := v.(*ListNode)
	*n = append(*n, node)
}

func (n *NodeSlice) Pop() interface{} {
	node := (*n)[len(*n) - 1]
	*n = (*n)[: len(*n) - 1]
	return node
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := &NodeSlice{}
	heap.Init(h)

	for _, list := range lists {
		if list == nil {
			continue
		}
		heap.Push(h, list)
	}

	ans := &ListNode{}
	head := ans

	for h.Len() > 0 {
		tmp := heap.Pop(h).(*ListNode)
		head.Next = tmp
		if tmp.Next != nil {
			heap.Push(h, tmp.Next)
		}

		head = head.Next
	}
	return ans.Next
}
//leetcode submit region end(Prohibit modification and deletion)
