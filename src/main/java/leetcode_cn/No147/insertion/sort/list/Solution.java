package leetcode_cn.No147.insertion.sort.list;

/**
 * @author rancho
 * @date 2019/10/5
 */
public class Solution {

    public class ListNode {
        int val;
        ListNode next;
        ListNode(int x) { val = x; }
    }

    public ListNode insertionSortList(ListNode head) {
        if (head == null) {
            return null;
        }

        ListNode dummy = new ListNode(-1);

        while (head != null) {
            ListNode current = dummy;
            ListNode next = head.next;

            while (current .next != null && current.next.val < head.val) {
                current = current.next;
            }

            // 在current后插入next
            head.next = current.next;
            current.next = head;
            head = next;
        }
        return dummy.next;
    }
}
