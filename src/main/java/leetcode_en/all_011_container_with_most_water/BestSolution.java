package leetcode_en.all_011_container_with_most_water;

/**
 * @author rancho
 * @date 2019-05-05
 */
public class BestSolution {

    public int maxArea(int[] height) {
        int max = 0;
        int i = 0, j = height.length - 1;

        while (i < j) {
            if (height[i] < height[j]) {
                max = Math.max(max, height[i] * (j - i));
                i++;
            } else {
                max = Math.max(max, height[j] * (j - i));
                j--;
            }
        }

        return max;
    }

}
