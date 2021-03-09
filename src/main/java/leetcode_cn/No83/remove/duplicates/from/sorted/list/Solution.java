package leetcode_cn.No83.remove.duplicates.from.sorted.list;

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

    public ListNode deleteDuplicates(ListNode head) {
        ListNode curNode = head;
        while (curNode != null && curNode.next != null) {
            if (curNode.val == curNode.next.val) {
                ListNode tmp = curNode.next;
                while (tmp != null && curNode.val == tmp.val) {
                    tmp = tmp.next;
                }
                curNode.next = tmp;
            } else {
                curNode = curNode.next;
            }
        }
        return head;
    }
}
