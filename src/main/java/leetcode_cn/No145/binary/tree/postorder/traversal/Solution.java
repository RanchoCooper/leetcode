package leetcode_cn.No145.binary.tree.postorder.traversal;

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

    public List<Integer>  postorderTraversal(TreeNode root) {
        if (root == null) {
            return result;
        }

        if (root.left != null) {
            postorderTraversal(root.left);
        }
        if (root.right != null) {
            postorderTraversal(root.right);
        }
        result.add(root.val);
        return result;

    }

}
