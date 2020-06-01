package leetcode_en.all_015_three_sum;

import java.util.Arrays;
import java.util.LinkedList;
import java.util.List;

/**
 * @author rancho
 * @date 2019-05-27
 */
public class Solution {

    public List<List<Integer>> threeSum(int[] nums) {
        Arrays.sort(nums);

        List<List<Integer>> res = new LinkedList<>();

        if (nums.length < 3) {
            return res;
        }

        for (int i = 0; i < nums.length - 2; i++) {
            if (nums[i] != nums[i + 1]) {
                int lo = i + 1;
                int hi = nums.length - 1;
                int sum = 0 - nums[i];

                while(lo < hi) {
                    if (nums[lo] + nums[hi] == sum) {
                        res.add(Arrays.asList(nums[i], nums[lo], nums[hi]));
                        while (lo < hi && nums[lo] == nums[lo + 1]) {
                            lo++;
                        }
                        while (lo < hi && nums[hi] == nums[hi - 1]) {
                            hi--;
                        }
                    } else if (nums[lo] + nums[hi] < sum) {
                        lo++;
                    } else {
                        hi--;
                    }
                }
            }
        }
        return res;
    }

}
