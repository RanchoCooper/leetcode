package leetcode_cn.No144.binary.tree.preorder.traversal;

import java.util.ArrayList;
import java.util.List;

/**
 * @author rancho
 * @date 2019/10/2
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

    private List<Integer> result = new ArrayList<>();

    public List<Integer>  preorderTraversal(TreeNode root) {
        if (root == null) {
            return result;
        }

        result.add(root.val);
        if (root.left != null) {
            preorderTraversal(root.left);
        }
        if (root.right != null) {
            preorderTraversal(root.right);
        }
        return result;

    }

}
