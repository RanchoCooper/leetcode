package leetcode_cn.No71.simplify.path;

import java.util.Stack;

/**
 * @author rancho
 * @date 2019/10/23
 */
public class Solution {

    public String simplifyPath(String path) {
        String[] elements = path.split("/");
        Stack<String> stack = new Stack<>();

        for (String element : elements) {
            if ("..".equals(element) && !stack.isEmpty()) {
                stack.pop();
            }
            if (!".".equals(element) && !"".equals(element) && !"..".equals(element)) {
                stack.push(element);
            }
        }
        if (stack.isEmpty()) {
            return "/";
        }
        StringBuffer result = new StringBuffer();
        for (int i = 0; i < stack.size(); i++) {
            result.append("/" + stack.get(i));
        }
        return result.toString();
    }

}
