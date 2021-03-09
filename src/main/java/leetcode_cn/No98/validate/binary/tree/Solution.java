package leetcode_cn.No98.validate.binary.tree;

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

    public boolean helper(TreeNode root, Integer lower, Integer upper) {
        if (root == null) {
            return true;
        }
        if (lower != null && root.val <= lower) {
            return false;
        }
        if (upper != null && root.val >= upper) {
            return false;
        }

        if (!helper(root.left, lower, root.val)) {
            return false;
        }
        if (!helper(root.right, root.val, upper)) {
            return false;
        }
        return true;
    }

    public boolean isValidBST(TreeNode root) {
        return helper(root, null, null);
    }

}
