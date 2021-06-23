package main

import (
	"fmt"
	"strconv"
	"sync"
)

//
// 有效的算符包括 +、-、*、/ 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。 
//
// 
//
// 说明： 
//
// 
// 整数除法只保留整数部分。 
// 给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。 
// 
//
// 
//
// 示例 1： 
//
// 
//输入：tokens = ["2","1","+","3","*"]
//输出：9
//解释：该算式转化为常见的中缀算术表达式为：((2 + 1) * 3) = 9
// 
//
// 示例 2： 
//
// 
//输入：tokens = ["4","13","5","/","+"]
//输出：6
//解释：该算式转化为常见的中缀算术表达式为：(4 + (13 / 5)) = 6
// 
//
// 示例 3： 
//
// 
//输入：tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
//输出：22
//解释：
//该算式转化为常见的中缀算术表达式为：
//  ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
//= ((10 * (6 / (12 * -11))) + 17) + 5
//= ((10 * (6 / -132)) + 17) + 5
//= ((10 * 0) + 17) + 5
//= (0 + 17) + 5
//= 17 + 5
//= 22 
//
// 
//
// 提示： 
//
// 
// 1 <= tokens.length <= 104 
// tokens[i] 要么是一个算符（"+"、"-"、"*" 或 "/"），要么是一个在范围 [-200, 200] 内的整数 
// 
//
// 
//
// 逆波兰表达式： 
//
// 逆波兰表达式是一种后缀表达式，所谓后缀就是指算符写在后面。 
//
// 
// 平常使用的算式则是一种中缀表达式，如 ( 1 + 2 ) * ( 3 + 4 ) 。 
// 该算式的逆波兰表达式写法为 ( ( 1 2 + ) ( 3 4 + ) * ) 。 
// 
//
// 逆波兰表达式主要有以下两个优点： 
//
// 
// 去掉括号后表达式无歧义，上式即便写成 1 2 + 3 4 + * 也可以依据次序计算出正确结果。 
// 适合用栈操作运算：遇到数字则入栈；遇到算符则取出栈顶两个数字进行计算，并将结果压入栈中。 
// 
// Related Topics 栈 
// 👍 362 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
type Stack struct {
	lock sync.RWMutex
	slice []string
}

func (s *Stack) Size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return len(s.slice)
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack) Push(value string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.slice = append(s.slice, value)
}

func (s *Stack) Pop() error {
	if s.IsEmpty() {
		return fmt.Errorf("stack is empty")
	}

	s.lock.Lock()
	defer s.lock.Unlock()

	s.slice = s.slice[:s.Size() - 1]
	return nil
}

func (s *Stack) Top() string {
	if s.IsEmpty() {
		return ""
	}
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.slice[len(s.slice) - 1]
}

func calculate(l, r, op string) int {
	a, _ := strconv.ParseInt(l, 10, 64)
	b, _ := strconv.ParseInt(r, 10, 64)
	switch op {
	case "+":
		return int(a + b)
	case "-":
		return int(a - b)
	case "*":
		return int(a * b)
	case "/":
		return int(a / b)
	}
	return 0
}

func evalRPN(tokens []string) int {
	var s Stack
	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			b := s.Top()
			s.Pop()
			a := s.Top()
			s.Pop()
			s.Push(strconv.Itoa(calculate(a, b, token)))
		} else {
			s.Push(token)
		}
	}
	r, _ := strconv.ParseInt(s.Top(), 10, 64)
	return int(r)
}
//leetcode submit region end(Prohibit modification and deletion)
