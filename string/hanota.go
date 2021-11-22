package main
import "C"

// Easy - 汉诺塔问题
// 原地将A柱子上的盘子移动到C，柱子只有A,B,C三个，表示栈。
//(1) 每次只能移动一个盘子;
//(2) 盘子只能从柱子顶端滑出移到下一根柱子;
//(3) 盘子只能叠在比它大的盘子上。
// 用栈将所有盘子从第一根柱子移到最后一根柱子。

//A, B, C := []int{2, 1, 0}, []int{}, []int{}
// 思路：递归，当一个问题可以分解为跟他结构类似的多个子问题时，就会用到递归.
// 当n=1时，直接把盘子从A移到C
// 当n>1时：
// 	 把A的n-1个盘子移动到B
//   把A剩下的最大的盘子移动到C
// 	 把B中的n-1个盘子移动到C
// 递归结束条件是A和B都空了，也就移动完了。
func hanota(A []int, B []int, C []int) {
	n := len(A)
	//I := createStack(A)
	//J := createStack(B)
	//K := createStack(C)
	move(n, A,B,C)
}

func popArr(a []int){
	if len(a)==0{
		return
	}
	a=a[:len(a)-1]
}
func move(n int, A, B, C []int) {
	if n == 1 {
		//C.push(A.pop()) //注意A.pop()不等于A[n-1]，但是否等于A[len(A)-1]呢
		C=append(C,A[len(A)-1])
		if A==
		return
	}
	move(n-1, A, C, B) //A中的n-1个盘子移动到B
	C.push(A.pop())    // A中最大的盘子移动到C
	move(n-1, B, A, C) //B中的n-1个盘子移动到C
}
