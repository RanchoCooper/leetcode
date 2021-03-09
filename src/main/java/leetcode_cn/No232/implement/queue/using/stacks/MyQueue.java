package leetcode_cn.No232.implement.queue.using.stacks;

import java.util.Stack;

/**
 * @author rancho
 * @date 2019/10/2
 *
 * Your MyQueue object will be instantiated and called as such:
 * MyQueue obj = new MyQueue();
 * obj.push(x);
 * int param_2 = obj.pop();
 * int param_3 = obj.peek();
 * boolean param_4 = obj.empty();
 */
class MyQueue {
    private int front;
    private int trail;
    private Stack<Integer> s1;
    private Stack<Integer> s2;

    /** Initialize your data structure here. */
    public MyQueue() {
        this.front = 0;
        this.trail = 0;
        this.s1 = new Stack<Integer>();
        this.s2 = new Stack<Integer>();
    }

    /** Push element x to the back of queue. */
    public void push(int x) {
        if (s1.isEmpty()) {
            this.front = x;
            this.trail = x;
        }
        while (!s1.isEmpty()) {
            s2.push(s1.pop());
        }
        s2.push(x);
        while(!s2.isEmpty()) {
            s1.push(s2.pop());
        }
        this.trail = x;

    }

    /** Removes the element from in front of queue and returns that element. */
    public int pop() {
        int result = s1.pop();
        if (!s1.isEmpty()) {
            this.front = s1.peek();
        } else {
            this.front = 0;
        }
        return result;
    }

    /** Get the front element. */
    public int peek() {
        return this.front;

    }

    /** Returns whether the queue is empty. */
    public boolean empty() {
        return s1.isEmpty();
    }
}

