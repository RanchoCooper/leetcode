package leetcode_cn.No23.merge.k.sorted.lists;

import java.util.Comparator;
import java.util.PriorityQueue;

/**
 * @author rancho
 * @date 2019/10/4
 */
class Solution {

    public class ListNode {
        int val;
        ListNode next;
        ListNode(int x) { val = x; }

        ListNode(Integer[] nums) {
            ListNode curNode = this;
            curNode.val = nums[0];
            for (int i = 1; i < nums.length; i++) {
                curNode.next = new ListNode(nums[i]);
                curNode = curNode.next;
            }
        }
    }

    public ListNode mergeKLists(ListNode[] lists) {
        if (lists == null || lists.length == 0) {
            return null;
        }

        PriorityQueue<ListNode> priorityQueue = new PriorityQueue<>(
                lists.length,
                Comparator.comparingInt(a -> a.val)
        );
        ListNode dummyNode = new ListNode(-1);
        ListNode curNode = dummyNode;

        // 所有列表首元素入队
        for (ListNode list : lists) {
            if (list != null) {
                priorityQueue.add(list);
            }
        }

        while(!priorityQueue.isEmpty()) {
            ListNode min = priorityQueue.poll();
            curNode.next = min;
            curNode = curNode.next;
            if (min.next != null) {
                priorityQueue.add(min.next);
            }
        }

        return dummyNode.next;
    }
}