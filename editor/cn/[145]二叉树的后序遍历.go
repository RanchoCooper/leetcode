package main
//给定一个二叉树，返回它的 后序 遍历。 
//
// 示例: 
//
// 输入: [1,null,2,3]  
//   1
//    \
//     2
//    /
//   3 
//
//输出: [3,2,1] 
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？ 
// Related Topics 栈 树 深度优先搜索 二叉树 
// 👍 615 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	postorder(root, &ans)
	return ans
}

func postorder(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	postorder(root.Left, ans)
	postorder(root.Right, ans)
	*ans = append(*ans, root.Val)

}
//leetcode submit region end(Prohibit modification and deletion)
