package leetcode_cn.No92.reverse.linked.list.between;

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

    public ListNode reverseBetween(ListNode head, int m, int n) {
        if (head == null || m > n) {
            return null;
        }
        ListNode curr = head;
        ListNode prev = null;
        while (m > 1) {
            prev = curr;
            curr = curr.next;
            m--;
            n--;
        }

        ListNode con = prev;
        ListNode tail = curr;
        ListNode next = null;

        while (n-- > 0) {
            next = curr.next;
            curr.next = prev;
            prev = curr;
            curr = next;
        }

        if (con != null) {
            con.next = prev;
        } else {
            head = prev;
        }

        tail.next = curr;

        return head;
    }
}
