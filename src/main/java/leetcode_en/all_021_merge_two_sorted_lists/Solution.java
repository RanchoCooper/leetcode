package leetcode_en.all_021_merge_two_sorted_lists;

/**
 * @author rancho
 * @date 2019-05-05
 */
public class Solution {

    class ListNode {

        int val;
        ListNode next;

        ListNode(int value) {
            this.val = value;
        }
    }

    public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
        if (l1 == null) {
            return l2;
        }

        if (l2 == null) {
            return l1;

        }

        if (l1.val < l2.val) {
            l1.next = mergeTwoLists(l1.next, l2);
            return l1;
        } else {
            l2.next = mergeTwoLists(l1, l2.next);
            return l2;
        }

    }

}
