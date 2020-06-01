package leetcode_cn.No61.rotate.list;

/**
 * @author rancho
 * @date 2019/10/4
 */
public class Solution {

    public class ListNode {
        int val;
        ListNode next;
        ListNode(int x) { val = x; }
   }

   public ListNode rotateRight(ListNode head, int k) {
        if (head == null) {
            return null;
        }
        if (head.next == null) {
            return head;
        }

        int n = 1;

        // 收尾相连

        ListNode oldTail = head;
        while(oldTail.next != null) {
            oldTail = oldTail.next;
            n++;
        }
        oldTail.next = head;

        ListNode newTail = head;
        for (int i = 0; i < n - k % n - 1; i++) {
            newTail = newTail.next;
        }
        ListNode newHead = newTail.next;
        newTail.next = null;
        return newHead;
   }

}
