package leetcode_cn.No155.min.stack;

import java.util.Stack;

/**
 * @author rancho
 * @date 2019/10/24
 */
public class MinStack {

    private Stack<Integer> data;
    private Stack<Integer> helper;
    public MinStack() {
        data = new Stack<>();
        helper = new Stack<>();
    }

    public void push(int x) {
        data.push(x);

        if (helper.isEmpty() || x <= helper.peek()) {
            helper.push(x);
        } else {
            helper.push(helper.peek());
        }
    }

    public void pop() {
        if (!data.isEmpty()) {
            data.pop();
            helper.pop();
        }
    }

    public int top() {
        if (!data.isEmpty()) {
            return data.peek();
        }
        throw new Error();
    }

    public int getMin() {
        if (!helper.isEmpty()) {
            return helper.peek();
        }
        throw new Error();
    }

}
