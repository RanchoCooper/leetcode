package leetcode_cn.No109.convert.sorted.list.to.binary.search.tree;

/**
 * @author rancho
 * @date 2019/10/8
 */
public class Solution {

    public static class ListNode {
        int val;
        ListNode next;
        ListNode(int x) { val = x; }
    }

    public static class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;
        TreeNode(int val) {
            this.val = val;
        }
    }

    static ListNode findMiddleElement(ListNode head) {
        if (head == null) {
            return null;
        }
        ListNode midPrev = null;
        ListNode slow = head;
        ListNode fast = head;

        while (fast != null && fast.next != null) {
            midPrev = slow;
            slow = slow.next;
            fast = fast.next.next;
        }
        if (midPrev != null) {
            // important!
            midPrev.next = null;
        }
        return slow;
    }


    public static TreeNode sortedListToBST(ListNode head) {
        if (head == null) {
            return null;
        }
        if (head.next == null) {
            return new TreeNode(head.val);
        }

        ListNode mid = findMiddleElement(head);
        TreeNode root = new TreeNode(mid.val);
        if (head == mid) {
            return root;
        }
        root.left = sortedListToBST(head);
        root.right = sortedListToBST(mid.next);
        return root;
    }

}
