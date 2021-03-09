package leetcode_cn.No20.valid.parentheses;

import java.util.HashMap;
import java.util.Stack;

/**
 * @author rancho
 * @date 2019/10/2
 */
public class Solution {

    private HashMap<Character, Character> mappings;

    public Solution() {
        this.mappings = new HashMap<Character, Character>();
        this.mappings.put(')', '(');
        this.mappings.put('}', '{');
        this.mappings.put(']', '[');
    }

    public boolean isValid(String s) {

        Stack<Character> stack = new Stack<Character>();

        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            if (!this.mappings.containsKey(c)) {
                stack.push(c);
            } else {
                char topElement = stack.isEmpty()? '#' : stack.pop();
                if (topElement != this.mappings.get(c)) {
                    return false;
                }
            }
        }
        return stack.isEmpty();

    }

}
