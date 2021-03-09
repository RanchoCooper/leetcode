package leetcode_cn.No150.evaluate.reverse.polish.notation;

import java.util.Stack;

/**
 * @author rancho
 * @date 2019/10/23
 */
public class Solution {


    public int evalRPN(String[] tokens) {
        Stack<Integer> stack = new Stack<>();

        for (String token : tokens) {
            if ("+".equals(token)) {
                stack.push(stack.pop() + stack.pop());
            } else if ("-".equals(token)) {
                stack.push(-stack.pop() + stack.pop());
            } else if ("*".equals(token)) {
                stack.push(stack.pop() * stack.pop());
            } else if ("/".equals(token)) {
                Integer right = stack.pop();
                stack.push(stack.pop() / right);
            } else {
                stack.push(Integer.parseInt(token));
            }
        }
        return stack.pop();
    }

}
