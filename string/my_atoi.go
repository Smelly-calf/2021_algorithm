package main

import "fmt"

// Medium - 将字符串转换成一个 32 位有符号整数

// 类型：整数运算，注意溢出。

// 这里我们要进行模式识别，一旦涉及整数的运算，我们需要注意溢出。
// 本题可能产生溢出的步骤在于推入、乘以 10 操作和累加操作都可能造成溢出。
// 对于溢出的处理方式通常可以转换为 INT_MAX 的逆操作。比如判断某数乘以 10 是否会溢出，那么就把该数和 INT_MAX 除以 10 进行比较。

// 建立确定有限状态机（deterministic finite automaton, DFA）：我们的程序在每个时刻有一个状态 s，每次从序列中输入一个字符 c，并根据字符 c 转移到下一个状态 s'。
// 这样，我们只需要建立一个覆盖所有情况的从 s 与 c 映射到 s' 的表格即可解决题目中的问题。
/*
|           | ' '   | +/-  | number     | other |
|  ----     | ---   | ---   | ---       | --- |
| start     | start | signed | in_number | end |
| signed    | end   | end   | in_number  | end  |
| in_number | end   | end   | in_number  | end |
| end       | end   | end   | end       | end |
*/
const (
	IntMax = 2<<30 - 1
	IntMin = -2 << 30
)

// 定义四种状态
const (
	start = iota
	signed
	inNumber
	end
)

// Automaton自动机
type Automaton struct {
	state int
	ans   int
	sign  int
}

// 有限状态机: s表示起始状态，c表示输入序列 => 返回target
func (*Automaton) machine(s int, c uint8) int {
	var s1 int
	if s == end {
		s1 = end
	}
	if s == start {
		if c == ' ' {
			s1 = start
		} else if c == '+' || c == '-' {
			s1 = signed
		} else if c >= '0' && c <= '9' {
			s1 = inNumber
		} else {
			s1 = end
		}
	}
	if s == signed || s == inNumber {
		if c == ' ' || c == '+' || c == '-' {
			s1 = end
		} else if c >= '0' && c <= '9' {
			s1 = inNumber
		} else {
			s1 = end
		}
	}
	return s1
}

func myAtoi(s string) int {
	// 初始化状态机
	am := Automaton{sign: 1}
	for i := 0; i < len(s); i++ {
		c := s[i]
		am.state = am.machine(am.state, c)
		if am.state == signed { // positive
			if s[i] == '+' {
				am.sign = 1
			} else {
				am.sign = -1
			}
		}
		if am.state == inNumber { // 推入s[i]
			if am.ans > IntMax/10 || am.ans < IntMin/10 {
				if am.sign == 1 {
					am.ans = IntMax
				} else {
					am.ans = IntMin
				}
				return am.ans
			}
			num := int(c - '0')
			am.ans = am.ans*10 + num
		}
	}
	fmt.Println("am.sign:", am.sign)
	fmt.Println("am.ans:", am.ans)
	return am.sign * am.ans
}

func main() {
	s := "42"               // 预期: 42
	s1 := "   -42"          // -42
	s2 := "4193 with words" // 4193
	s3 := "words and 987"   // 0
	s4 := "-91283472332"    //  -2^31 = -2147483648
	fmt.Println(myAtoi(s))
	fmt.Println(myAtoi(s1))
	fmt.Println(myAtoi(s2))
	fmt.Println(myAtoi(s3))
	fmt.Println(myAtoi(s4))
}

