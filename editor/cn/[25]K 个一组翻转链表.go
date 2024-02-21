//给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。 
//
// k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。 
//
// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。 
//
// 
//
// 示例 1： 
// 
// 
//输入：head = [1,2,3,4,5], k = 2
//输出：[2,1,4,3,5]
// 
//
// 示例 2： 
//
// 
//
// 
//输入：head = [1,2,3,4,5], k = 3
//输出：[3,2,1,4,5]
// 
//
// 
//提示：
//
// 
// 链表中的节点数目为 n 
// 1 <= k <= n <= 5000 
// 0 <= Node.val <= 1000 
// 
//
// 
//
// 进阶：你可以设计一个只用 O(1) 额外内存空间的算法解决此问题吗？ 
//
// 
// 
//
// Related Topics 递归 链表 👍 2280 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	// 分组(找到每一组的开始、结尾)，按组遍历
	// 保护节点
	protect := &ListNode{Val: 0, Next: head}
	last := protect

	for head != nil {
		tail := getTail(head, k)
		if tail == nil {
			break
		}

		nextGroup := tail.Next
		// 处理head到tail之间的k-1条边的反转
		reverseList(head, tail)
		// 上一组跟本组的新head(旧tail)建立联系
		last.Next = tail
		// 本组的新结尾(head)跟下一组建立联系
		head.Next = nextGroup

		// 分组遍历
		last = head
		head = nextGroup
	}
	return protect.Next
}

func getTail(head *ListNode, k int) *ListNode {
	for head != nil {
		k--
		if k == 0 {
			break
		}
		head = head.Next
	}
	return head
}

func reverseList(head, tail *ListNode) {
	if head == tail {
		return
	}
	last := head
	head = head.Next
	for head != tail {
		next := head.Next
		head.Next = last
		last = head
		head = next
	}
	tail.Next = last
}
//leetcode submit region tail(Prohibit modification and deletion)
