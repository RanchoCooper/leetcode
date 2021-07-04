package main
//给定一个非空且只包含非负数的整数数组 nums，数组的度的定义是指数组里任一元素出现频数的最大值。 
//
// 你的任务是在 nums 中找到与 nums 拥有相同大小的度的最短连续子数组，返回其长度。 
//
// 
//
// 示例 1： 
//
// 
//输入：[1, 2, 2, 3, 1]
//输出：2
//解释：
//输入数组的度是2，因为元素1和2的出现频数最大，均为2.
//连续子数组里面拥有相同度的有如下所示:
//[1, 2, 2, 3, 1], [1, 2, 2, 3], [2, 2, 3, 1], [1, 2, 2], [2, 2, 3], [2, 2]
//最短连续子数组[2, 2]的长度为2，所以返回2.
// 
//
// 示例 2： 
//
// 
//输入：[1,2,2,3,1,4,2]
//输出：6
// 
//
// 
//
// 提示： 
//
// 
// nums.length 在1到 50,000 区间范围内。 
// nums[i] 是一个在 0 到 49,999 范围内的整数。 
// 
// Related Topics 数组 哈希表 
// 👍 351 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
import "math"

func findShortestSubArray(nums []int) int {
	count := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		count[nums[i]] += 1
	}

	degree, values := 0, make([]int, 0)
	for _, v := range count {
		degree = max(degree, v)
	}
	for k, v := range count {
		if v == degree {
			values = append(values, k)
		}
	}

	length := math.MaxInt32
	for _, value := range values {
		left, right := 0, len(nums) - 1
		for ; left < len(nums); left++ {
			if nums[left] == value {
				break
			}
		}
		for ; right >= 0; right-- {
			if nums[right] == value {
				break
			}
		}
		length = min(length, right - left + 1)
	}
	return length

}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
//leetcode submit region end(Prohibit modification and deletion)
