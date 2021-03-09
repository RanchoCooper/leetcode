package leetcode_en.all_019_remove_nth_node_from_end_of_list;

/**
 * @author rancho
 * @date 2019-05-05
 */
public class Solution {

    class ListNode {

        int value;
        ListNode next;

        ListNode(int value) {
            this.value = value;
        }
    }

    public ListNode removeNthFromEnd(ListNode head, int n) {

        ListNode dummy = new ListNode(-1);

        dummy.next = head;
        ListNode first = dummy, second = dummy;

        for (int i = 0; i <= n ; i ++) {
            first = first.next;
        }

        while (first != null) {
            first = first.next;
            second = second.next;
        }

        second.next = second.next.next;
        return dummy.next;

    }

}
