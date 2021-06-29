package main

import (
	"sort"
)

//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重
//复的三元组。 
//
// 注意：答案中不可以包含重复的三元组。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
// 
//
// 示例 2： 
//
// 
//输入：nums = []
//输出：[]
// 
//
// 示例 3： 
//
// 
//输入：nums = [0]
//输出：[]
// 
//
// 
//
// 提示： 
//
// 
// 0 <= nums.length <= 3000 
// -105 <= nums[i] <= 105 
// 
// Related Topics 数组 双指针 排序 
// 👍 3453 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	sort.Sort(sort.IntSlice(nums))

	ans := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}
		jks := myTwoSum(nums, i + 1, -nums[i])
		for _, jk := range jks {
			ans = append(ans, []int{nums[i], jk[0], jk[1]})
		}
	}
	return ans
}

func myTwoSum(nums []int, start, target int) [][]int {
	ans := make([][]int, 0)
	j := len(nums) - 1
	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i - 1] {
			continue
		}
		for ; i < j && nums[i] + nums[j] > target; j-- {

		}
		if i < j && nums[i] + nums[j] == target {
			ans = append(ans, []int{nums[i], nums[j]})
		}
	}
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
