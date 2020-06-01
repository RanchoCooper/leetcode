package leetcode_en.all_011_container_with_most_water;

/**
 * @author rancho
 * @date 2019-05-05
 */
public class Solution {

    public int maxArea(int[] height) {
        int left = 0, right = height.length - 1;
        int max = 0, temp = 0;

        while (left < right) {
            if (height[left] > height [right]) {
                temp = (right - left) * height[right];
                right--;
            } else {
                temp = (right - left) * height[left];
                left++;
            }

            max = max > temp ? max : temp;
        }

        return max;

    }

}
