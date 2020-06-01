package leetcode_cn.No19.remove.nth.node.from.end.of.list;

/**
 * @author rancho
 * @date 2019/10/4
 */
public class Solution {

    public class ListNode {
        int val;
        ListNode next;
        ListNode(int val) {
            this.val = val;
        }
    }

    public ListNode removeNthFromEnd(ListNode head, int n) {
        if (head == null || n <= 0) {
            return null;
        }

        ListNode dummy = new ListNode(0);
        dummy.next = head;
        ListNode first  = dummy;
        ListNode second = dummy;

        for (int i = 0; i < n + 1; i++) {
            first = first.next;
        }

        while(first != null) {
            first = first.next;
            second = second.next;
        }

        // delete node
        second.next = second.next.next;
        return dummy.next;
    }

}
