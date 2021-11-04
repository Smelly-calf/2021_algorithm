package main

// 将字符串转换成一个 32 位有符号整数

// 类型：整数运算，注意溢出。

// 这里我们要进行模式识别，一旦涉及整数的运算，我们需要注意溢出。
// 本题可能产生溢出的步骤在于推入、乘以 1010 操作和累加操作都可能造成溢出。
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