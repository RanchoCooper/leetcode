package leetcode_cn.No104.maximum.depth.of.binary.tree;

/**
 * @author rancho
 * @date 2019/10/7
 */
public class Solution {

    public class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;
        TreeNode(int val) {
            this.val = val;
        }
    }

    public int maxDepth(TreeNode root) {
        if (root == null) {
            return 0;
        }
        int left_hight = maxDepth(root.left);
        int right_hight = maxDepth(root.right);
        return Math.max(left_hight, right_hight) + 1;
    }

}
