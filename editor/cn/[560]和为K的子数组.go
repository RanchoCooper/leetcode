package main
//给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。 
//
// 示例 1 : 
//
// 
//输入:nums = [1,1,1], k = 2
//输出: 2 , [1,1] 与 [1,1] 为两种不同的情况。
// 
//
// 说明 : 
//
// 
// 数组的长度为 [1, 20,000]。 
// 数组中元素的范围是 [-1000, 1000] ，且整数 k 的范围是 [-1e7, 1e7]。 
// 
// Related Topics 数组 哈希表 
// 👍 960 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func subarraySum(nums []int, k int) int {
	sum := make([]int, len(nums)+1)

	for i := 0; i < len(nums); i++ {
		sum[i+1] = sum[i] + nums[i]
	}

	ans := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if sum[j+1] - sum[i] == k {
				ans++
			}
		}
	}
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
