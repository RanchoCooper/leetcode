package leetcode_cn.No215.kth.largest.element.in.an.array;

/**
 * @author rancho
 * @date 2019/10/21
 */
public class Solution {

    public int findKthLargest(int[] nums, int k) {
        if (nums == null || nums.length < k) {
            return 0;
        }
        return partition(nums, 0, nums.length - 1, k);
    }

    public int partition(int[] nums, int left, int right, int k) {
        int start = left;
        int end = right;
        int key = nums[start];

        while (start < end) {
            while (start < end && nums[end] <= key) {
                end--;
            }
            nums[start] = nums[end];
            while (start < end && nums[start] > key) {
                start++;
            }
            nums[end] = nums[start];
        }
        nums[start] = key;
        if (start < k - 1) {
            return partition(nums, start + 1, right, k);
        }
        if (start > k - 1) {
            return partition(nums, left, start - 1, k);
        }
        return key;
    }

}
