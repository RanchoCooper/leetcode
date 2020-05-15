//Given a 32-bit signed integer, reverse digits of an integer. 
//
// Example 1: 
//
// 
//Input: 123
//Output: 321
// 
//
// Example 2: 
//
// 
//Input: -123
//Output: -321
// 
//
// Example 3: 
//
// 
//Input: 120
//Output: 21
// 
//
// Note: 
//Assume we are dealing with an environment which could only store integers with
//in the 32-bit signed integer range: [−231, 231 − 1]. For the purpose of this pro
//blem, assume that your function returns 0 when the reversed integer overflows. 
// Related Topics 数学


//leetcode submit region begin(Prohibit modification and deletion)
package main
func reverse(x int) int {
	var y = 0
	for x != 0 {
		y = y * 10 + x % 10
		x /= 10
	}
	if y > (1 << 31 - 1) || y < (-1 << 31) {
		return 0
	}
	return y
}

//leetcode submit region end(Prohibit modification and deletion)
