package main
//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。 
//
// 有效字符串需满足： 
//
// 
// 左括号必须用相同类型的右括号闭合。 
// 左括号必须以正确的顺序闭合。 
// 
//
// 
//
// 示例 1： 
//
// 
//输入：s = "()"
//输出：true
// 
//
// 示例 2： 
//
// 
//输入：s = "()[]{}"
//输出：true
// 
//
// 示例 3： 
//
// 
//输入：s = "(]"
//输出：false
// 
//
// 示例 4： 
//
// 
//输入：s = "([)]"
//输出：false
// 
//
// 示例 5： 
//
// 
//输入：s = "{[]}"
//输出：true 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 104 
// s 仅由括号 '()[]{}' 组成 
// 
// Related Topics 栈 字符串 
// 👍 2460 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	m := map[string]string {
		")": "(",
		"]": "[",
		"}": "{",
	}
	var stack []string
	for _, c := range s {
		c := string(c)
		if _, ok := m[c]; ok {
			if len(stack) > 0 {
				if stack[len(stack) -1] == m[c] {
					stack = stack[:len(stack) - 1]
				} else {
					return false
				}
			} else {
				return false
			}
		} else {
			stack = append(stack, c)
		}
	}
	return len(stack) == 0
}
//leetcode submit region end(Prohibit modification and deletion)
