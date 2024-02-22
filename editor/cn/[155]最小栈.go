//设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。 
//
// 实现 MinStack 类: 
//
// 
// MinStack() 初始化堆栈对象。 
// void push(int val) 将元素val推入堆栈。 
// void pop() 删除堆栈顶部的元素。 
// int top() 获取堆栈顶部的元素。 
// int getMin() 获取堆栈中的最小元素。 
// 
//
// 
//
// 示例 1: 
//
// 
//输入：
//["MinStack","push","push","push","getMin","pop","top","getMin"]
//[[],[-2],[0],[-3],[],[],[],[]]
//
//输出：
//[null,null,null,null,-3,null,0,-2]
//
//解释：
//MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.getMin();   --> 返回 -3.
//minStack.pop();
//minStack.top();      --> 返回 0.
//minStack.getMin();   --> 返回 -2.
// 
//
// 
//
// 提示： 
//
// 
// -2³¹ <= val <= 2³¹ - 1 
// pop、top 和 getMin 操作总是在 非空栈 上调用 
// push, pop, top, and getMin最多被调用 3 * 10⁴ 次 
// 
//
// Related Topics 栈 设计 👍 1727 👎 0

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
type MinStack struct {
	stack []int
	min int
}


func Constructor() MinStack {
	return MinStack{
		stack: make([]int, 0),
		min: math.MaxInt32,
	}
}


func (this *MinStack) Push(val int)  {
	this.stack = append(this.stack, val)
	if val < this.min {
		this.min = val
	}
}


func (this *MinStack) Pop()  {
	last := this.Top()
	this.stack = this.stack[:len(this.stack) - 1]

	if last == this.min {
		min := math.MaxInt32
		for _, v := range this.stack {
			if v < min {
				min = v
			}
		}
		this.min = min
	}
}


func (this *MinStack) Top() int {
	return this.stack[len(this.stack) - 1]
}


func (this *MinStack) GetMin() int {
	return this.min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
//leetcode submit region end(Prohibit modification and deletion)
