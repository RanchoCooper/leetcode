package main

import (
	"sort"
	"strconv"
	"strings"
)

//给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。
//
// 注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [10,2]
//输出："210" 
//
// 示例 2： 
//
// 
//输入：nums = [3,30,34,5,9]
//输出："9534330"
// 
//
// 示例 3： 
//
// 
//输入：nums = [1]
//输出："1"
// 
//
// 示例 4： 
//
// 
//输入：nums = [10]
//输出："10"
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 100 
// 0 <= nums[i] <= 109 
// 
// Related Topics 贪心 字符串 排序 
// 👍 757 👎 0

type SliceInt []int

func (s SliceInt) Len() int {
	return len(s)
}

func (s SliceInt) Less(i, j int) bool {
	a, b := strconv.Itoa(s[i]), strconv.Itoa(s[j])
	return a + b < b + a
}

func (s SliceInt) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

//leetcode submit region begin(Prohibit modification and deletion)
func largestNumber(nums []int) string {
	var ns SliceInt
	ns = nums
	sort.Sort(ns)

	s := make([]string, 0)
	for i := len(ns); i > 0; i-- {
		tmp := strconv.Itoa(nums[i - 1])
		s = append(s, tmp)
	}
	ans := strings.Join(s, "")
	if ans[0] == '0' {
		return "0"
	}
	return ans

}
//leetcode submit region end(Prohibit modification and deletion)
