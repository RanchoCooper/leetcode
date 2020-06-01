package leetcode_cn.No225.implement.stack.using.queues;

import java.util.LinkedList;
import java.util.Queue;

/**
 * @author rancho
 * @date 2019/10/25
 */
public class MyStack {

    private Queue<Integer> q1 = new LinkedList<>();
    private Queue<Integer> q2 = new LinkedList<>();
    private int top;

    public MyStack() {

    }

    public void push(int x) {
        q1.add(x);
        top = x;
    }

    public int pop() {
        while(q1.size() > 1) {
            top = q1.remove();
            q2.add(top);
        }
        int result = q1.remove();

        Queue<Integer> temp = q1;
        q1 = q2;
        q2 = temp;
        return result;
    }

    public int top() {
        return top;
    }

    public boolean empty() {
        return q1.isEmpty();
    }


}
