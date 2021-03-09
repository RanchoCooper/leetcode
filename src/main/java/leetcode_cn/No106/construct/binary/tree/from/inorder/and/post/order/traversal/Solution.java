package leetcode_cn.No106.construct.binary.tree.from.inorder.and.post.order.traversal;

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

    public TreeNode buildTree(int[] inorder, int[] postorder) {
        int inLen = inorder.length;
        int postLen = postorder.length;
        if (inLen != postLen) {
            return null;
        }
        if (inorder.length != postorder.length) {
            return null;
        }
        return buildTree(
                inorder, 0, inLen - 1,
                postorder, 0, postLen - 1
                );
    }

    private TreeNode buildTree(int[] inorder, int inLeft, int inRight,
                               int[] postorder, int postLeft, int postRight) {
        if (inLeft > inRight || postLeft > postRight) {
            return null;
        }
        int pivot = postorder[postRight];
        int pivotIndex = inLeft;
        while (inorder[pivotIndex] != pivot) {
            pivotIndex++;
        }
        TreeNode root = new TreeNode(pivot);
        root.left = buildTree(
                inorder, inLeft, pivotIndex - 1,
                // postorder, postLeft, pivotIndex - inRight + postRight - 1
                postorder, postLeft, pivotIndex - inLeft + postLeft - 1
        );
        root.right = buildTree(
                inorder, pivotIndex + 1, inRight,
                // postorder, pivotIndex - inRight + postRight, postRight - 1
                postorder, pivotIndex - inLeft + postLeft, postRight -1
        );
        return root;
    }

}
